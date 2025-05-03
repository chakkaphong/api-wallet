package models

import (
	"api-wallet/internal/entities"
)

type ConfirmTransactionRequest struct {
	TransactionId string `json:"transaction_id" validate:"required"`
}

type ConfirmTransactionResponse struct {
	TransactionId string                         `json:"transaction_id"`
	UserId        int                            `json:"user_id"`
	Amount        float64                        `json:"amount"`
	Status        entities.TransactionStatusEnum `json:"status"`
	Balance       float64                        `json:"balance"`
}
