package wallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"go-web3-wallet-backend/contracts/mininft"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func MintNFT(to string) (string, error) {
	rpc := viper.GetString("chain.rpc_url")
	pkHex := viper.GetString("chain.private_key")
	if pkHex == "" {
		// 如果配置文件中没有，从环境变量读取
		pkHex = os.Getenv("CHAIN_PRIVATE_KEY")
	}
	if pkHex == "" {
		return "", errors.New("private key not configured: set CHAIN_PRIVATE_KEY environment variable or chain.private_key in config")
	}
	chainID := big.NewInt(viper.GetInt64("chain.chain_id"))
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // 不转 ETH
	auth.GasLimit = uint64(300000) // NFT mint 足够
	auth.GasPrice = gasPrice

	contractAddr := common.HexToAddress(MiniNFTAddress)
	nft, err := mininft.NewMiniNFT(contractAddr, client)
	if err != nil {
		return "", err
	}

	tx, err := nft.Mint(auth, common.HexToAddress(to))
	if err != nil {
		return "", err
	}

	log.Println("Mint tx sent:", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}
