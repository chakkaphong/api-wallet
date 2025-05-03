package wallets

import (
	"api-wallet/internal/api/wallets/models"
	"api-wallet/internal/entities"
	"context"
)

type Service interface {
	VerifyTransaction(ctx context.Context, payload models.VerifyRequest) (*entities.Transaction, error)
	ConfirmTransaction(ctx context.Context, payload models.ConfirmTransactionRequest) (*models.ConfirmTransactionResponse, error)
}

type Repository interface {
	GetWalletByUserId(ctx context.Context, userId int) (*entities.Wallet, error)
}
