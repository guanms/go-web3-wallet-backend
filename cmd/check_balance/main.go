package main

import (
	"fmt"
	"os"

	"go-web3-wallet-backend/internal/wallet"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run cmd/check_balance/main.go <以太坊地址>")
		fmt.Println("示例: go run cmd/check_balance/main.go 0xYourAddress")
		os.Exit(1)
	}

	address := os.Args[1]
	fmt.Printf("正在查询地址 %s 的 NFT 余额...\n\n", address)

	wallet.PrintNFTBalance(address)

	fmt.Println("\n=== 提示 ===")
	fmt.Println("1. 确保 MetaMask 连接到 Sepolia 测试网")
	fmt.Println("2. 在 MetaMask 中导入 NFT 合约地址：")
	fmt.Printf("   合约地址: %s\n", wallet.MiniNFTAddress)
	fmt.Println("3. Token ID: 从 0 开始（1, 2, 3...）")
	fmt.Println("4. Token URI: 可以在合约浏览器中查看")
	fmt.Println("\nSepolia 测试网浏览器：")
	fmt.Printf("https://sepolia.etherscan.io/address/%s\n", wallet.MiniNFTAddress)
}
