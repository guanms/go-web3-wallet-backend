package server

import (
	"go-web3-wallet-backend/internal/wallet"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	api.POST("/nft/mint", func(c *gin.Context) {
		var req struct {
			To string `json:"to"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		txHash, err := wallet.MintNFT(req.To)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"txHash": txHash,
		})
	})
}
