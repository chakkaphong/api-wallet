package users

import (
	"api-wallet/internal/entities"
	"context"
)

type Service interface{}

type Repository interface {
	GetUserByID(ctx context.Context, userID int) (*entities.User, error)
}
