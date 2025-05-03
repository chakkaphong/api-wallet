package transactions

import (
	"api-wallet/internal/entities"
	"context"
)

type Service interface{}

type Repository interface {
	SaveTransaction(ctx context.Context, payload entities.Transaction) (*entities.Transaction, error)
	GetTransactionById(ctx context.Context, transactionId string) (*entities.Transaction, error)
	ConfirmTransaction(ctx context.Context, transaction entities.Transaction) error
}
