package repo

import (
	"api-wallet/internal/entities"
	"context"
)

func (r *UserRepository) GetUserByID(ctx context.Context, userID int) (*entities.User, error) {
	var user entities.User

	if err := r.postgres.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
