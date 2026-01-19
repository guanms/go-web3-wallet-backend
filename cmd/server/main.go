package main

import (
	"go-web3-wallet-backend/internal/server"
	"go-web3-wallet-backend/internal/wallet"
	"log"

	_ "github.com/joho/godotenv/autoload" // 自动加载 .env 文件
)

func main() {
	wallet.TestReadNFT()

	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}
