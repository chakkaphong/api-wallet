package repo

import (
	"api-wallet/internal/entities"
	"context"
)

func (r *WalletRepository) GetWalletByUserId(ctx context.Context, userId int) (*entities.Wallet, error) {
	var wallet entities.Wallet

	if err := r.postgres.WithContext(ctx).
		Where("user_id = ?", userId).
		First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
