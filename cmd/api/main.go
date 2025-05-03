package main

import (
	"api-wallet/internal/global/logger"
	"api-wallet/internal/server"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	log := logger.Init()
	defer log.Sync()

	s := server.NerServer()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Listen error: %v", zap.String("error", err.Error()))
	}
}
