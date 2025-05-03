package repo

import (
	"api-wallet/internal/entities"
	"context"
)

func (r *TransactionRepository) GetTransactionById(ctx context.Context, transactionId string) (*entities.Transaction, error) {
	var transaction entities.Transaction

	if err := r.postgres.
		WithContext(ctx).
		Where("transaction_id = ?", transactionId).
		Preload("Wallet").
		First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
