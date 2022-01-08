package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/serg-pe/blockchainrest/pkg/logging"
)

type Server struct {
	config *ServerConfig
	router *gin.Engine
}

func NewServer(configPath string) *Server {
	config := CreateConfigFromFile(configPath)
	server := Server{
		config: config,
		router: NewRouter(),
	}
	return &server
}

func (server Server) Start() {
	err := server.router.Run(server.config.ServerAddress)
	if err != nil {
		logging.GetLogger().Fatal("Server error")
	}
}
