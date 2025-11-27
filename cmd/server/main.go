package main

import (
	"go-web3-wallet-backend/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
