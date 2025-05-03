package models

import (
	"api-wallet/internal/entities"
	"time"
)

type VerifyRequest struct {
	UserId        int                        `json:"user_id" validate:"required"`
	Amount        float64                    `json:"amount" validate:"required,gt=0"`
	PaymentMethod entities.PaymentMethodEnum `json:"payment_method" validate:"required,oneof=credit_card"`
}

type VerifyResponse struct {
	TransactionId string                         `json:"transaction_id"`
	UserId        int                            `json:"user_id"`
	Amount        float64                        `json:"amount"`
	PaymentMethod entities.PaymentMethodEnum     `json:"payment_method"`
	Status        entities.TransactionStatusEnum `json:"status"`
	ExpiresAt     time.Time                      `json:"expires_at"`
}
