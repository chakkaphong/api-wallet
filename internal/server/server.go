package server

import (
	"api-wallet/configs"
	"api-wallet/internal/database"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port     int
	postgres database.Service
}

func NerServer() *http.Server {
	cfg := configs.GetConfig()

	postgres := database.New()

	NewServer := &Server{
		port:     cfg.App.Port,
		postgres: postgres,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  cfg.Server.IdleTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.ReadTimeout,
	}

	log.Default().Printf("Server is running, port:%d", NewServer.port)

	return server
}
