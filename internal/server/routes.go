package server

import (
	"api-wallet/configs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	transactionRepository "api-wallet/internal/api/transactions/repo"
	userRepository "api-wallet/internal/api/users/repo"

	walletRepository "api-wallet/internal/api/wallets/repo"
	walletRoutes "api-wallet/internal/api/wallets/routes"
	walletService "api-wallet/internal/api/wallets/services"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (s *Server) RegisterRoutes() *echo.Echo {
	cfg := configs.GetConfig()
	e := echo.New()
	var apiRoot echoswagger.ApiRoot

	apiRoot = echoswagger.New(e, "docs", &echoswagger.Info{
		Title:       cfg.App.Name,
		Description: "Template API Documentation",
		Version:     "0.0.1",
	})

	e.Validator = &CustomValidator{validator: validator.New()}

	// Repo
	walletRepo := walletRepository.NewWalletRepository(s.postgres)
	userRepo := userRepository.NewUserRepository(s.postgres)
	transactionRepo := transactionRepository.NewTransactionRepository(s.postgres)

	// Service
	walletSrv := walletService.NewService(walletRepo, userRepo, transactionRepo)

	// Routes
	walletRoutes.RegisterProductRoutes(apiRoot, walletSrv)

	return e
}
