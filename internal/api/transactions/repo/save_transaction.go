package repo

import (
	"api-wallet/internal/entities"
	"context"
)

func (r *TransactionRepository) SaveTransaction(ctx context.Context, payload entities.Transaction) (*entities.Transaction, error) {
	tx := r.postgres.WithContext(ctx)

	if err := tx.Create(&payload).Error; err != nil {
		return nil, err
	}

	return &payload, nil
}
