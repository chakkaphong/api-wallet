package services

import (
	"api-wallet/internal/api/transactions"
	"api-wallet/internal/api/users"
	"api-wallet/internal/api/wallets"
)

type WalletServices struct {
	repo            wallets.Repository
	userRepo        users.Repository
	transactionRepo transactions.Repository
}

func NewService(
	repo wallets.Repository,
	userRepo users.Repository,
	transactionRepo transactions.Repository,
) wallets.Service {
	return &WalletServices{
		repo:            repo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}
