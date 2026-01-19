package main

import (
	"fmt"
	"log"
	"os"

	"go-web3-wallet-backend/internal/wallet"
)

func main() {
	// 从命令行参数获取交易 hash
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run cmd/check_tx/main.go <交易哈希>")
		fmt.Println("示例: go run cmd/check_tx/main.go 0xbe2d5097b1b6879c91df3cbf28d13d28a247e6d7e55d2813f57f8e8a9a938f66")
		os.Exit(1)
	}

	txHash := os.Args[1]
	fmt.Printf("正在检查交易: %s\n\n", txHash)

	// 检查交易状态
	err := wallet.CheckTransactionStatus(txHash)
	if err != nil {
		log.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ 交易检查完成")
}
