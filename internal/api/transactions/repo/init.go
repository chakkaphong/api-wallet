package repo

import (
	"api-wallet/internal/api/transactions"
	"api-wallet/internal/database"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	postgres *gorm.DB
}

func NewTransactionRepository(postgres database.Service) transactions.Repository {
	return &TransactionRepository{
		postgres: postgres.DB(),
	}
}
