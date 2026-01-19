package wallet

import (
	"go-web3-wallet-backend/contracts/mininft"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const MiniNFTAddress = "0x01Bc43B9438a08509DDC6a138c6d909B3CbECa21"

func TestReadNFT() {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_sepolia/221dc68cf10676e52c6d352b7e55da4dee955be9f2b8d3cdcab008bb6c96ebd4")
	if err != nil {
		log.Fatal(err)
	}
	contractAddr := common.HexToAddress(MiniNFTAddress)
	nft, err := mininft.NewMiniNFT(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	nextId, err := nft.NextTokenId(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Next token id: %d\n", nextId)
}
