package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	loadConfig()

	engine := gin.Default()
	registerRoutes(engine)

	return &Server{engine: engine}
}

func (s *Server) Run() error {
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	return s.engine.Run(fmt.Sprintf(":%s", port))
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load config: " + err.Error())
	}
}
