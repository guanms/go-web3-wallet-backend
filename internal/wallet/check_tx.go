package wallet

import (
	"context"
	"fmt"
	"go-web3-wallet-backend/contracts/mininft"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CheckTransactionStatus 检查交易状态
func CheckTransactionStatus(txHash string) error {
	rpc := "https://rpc.ankr.com/eth_sepolia/221dc68cf10676e52c6d352b7e55da4dee955be9f2b8d3cdcab008bb6c96ebd4"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}

	txHashHex := common.HexToHash(txHash)

	// 1. 检查交易收据
	receipt, err := client.TransactionReceipt(context.Background(), txHashHex)
	if err != nil {
		return fmt.Errorf("获取交易收据失败: %v", err)
	}

	fmt.Println("=== 交易信息 ===")
	fmt.Printf("交易 Hash: %s\n", txHash)
	fmt.Printf("区块号: %d\n", receipt.BlockNumber)
	fmt.Printf("状态: %d (1=成功, 0=失败)\n", receipt.Status)
	fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
	fmt.Printf("Cumulative Gas Used: %d\n", receipt.CumulativeGasUsed)

	if receipt.Status == 0 {
		return fmt.Errorf("交易失败")
	}

	// 2. 检查发送者地址
	tx, _, err := client.TransactionByHash(context.Background(), txHashHex)
	if err != nil {
		return fmt.Errorf("获取交易详情失败: %v", err)
	}

	fmt.Printf("发送者地址: %s\n", tx.To().Hex())

	// 3. 检查 NFT 合约的下一个 Token ID
	contractAddr := common.HexToAddress(MiniNFTAddress)
	nft, err := mininft.NewMiniNFT(contractAddr, client)
	if err != nil {
		return fmt.Errorf("连接 NFT 合约失败: %v", err)
	}

	nextTokenId, err := nft.NextTokenId(nil)
	if err != nil {
		return fmt.Errorf("获取 NextTokenId 失败: %v", err)
	}

	fmt.Printf("\n=== NFT 合约信息 ===")
	fmt.Printf("合约地址: %s\n", MiniNFTAddress)
	fmt.Printf("下一个 Token ID: %s\n", nextTokenId.String())

	// 4. 检查指定地址的 NFT 余额
	// 这里需要知道 mint 到哪个地址了，可以手动查询
	fmt.Printf("\n=== 提示 ===\n")
	fmt.Printf("请在 Sepolia 测试网浏览器查看交易详情:\n")
	fmt.Printf("https://sepolia.etherscan.io/tx/%s\n", txHash)

	return nil
}

// GetNFTBalance 查询指定地址的 NFT 余额
func GetNFTBalance(address string) (*big.Int, error) {
	rpc := "https://rpc.ankr.com/eth_sepolia/221dc68cf10676e52c6d352b7e55da4dee955be9f2b8d3cdcab008bb6c96ebd4"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(MiniNFTAddress)
	nft, err := mininft.NewMiniNFT(contractAddr, client)
	if err != nil {
		return nil, err
	}

	balance, err := nft.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// PrintNFTBalance 打印 NFT 余额
func PrintNFTBalance(address string) {
	balance, err := GetNFTBalance(address)
	if err != nil {
		log.Printf("查询余额失败: %v\n", err)
		return
	}

	fmt.Printf("地址 %s 的 NFT 余额: %s\n", address, balance.String())
}
