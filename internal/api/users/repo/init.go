package repo

import (
	"api-wallet/internal/api/users"
	"api-wallet/internal/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	postgres *gorm.DB
}

func NewUserRepository(postgres database.Service) users.Repository {
	return &UserRepository{
		postgres: postgres.DB(),
	}
}
