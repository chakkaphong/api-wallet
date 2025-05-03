package repo

import (
	"api-wallet/internal/api/wallets"
	"api-wallet/internal/database"

	"gorm.io/gorm"
)

type WalletRepository struct {
	postgres *gorm.DB
}

func NewWalletRepository(postgres database.Service) wallets.Repository {
	return &WalletRepository{
		postgres: postgres.DB(),
	}
}
