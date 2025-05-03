package services

import (
	"api-wallet/internal/api/wallets/models"
	"api-wallet/internal/entities"
	"context"
	"errors"
	"net/http"
	"time"

	appError "api-wallet/internal/global/errors"
	"api-wallet/internal/global/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ErrTransactionNotFound  string = "transaction not found"
	ErrInvalidConfirmStatus string = "can confirm only verified transaction"
	ErrTransactionExpred    string = "transaction is already expired"
)

func (s *WalletServices) ConfirmTransaction(ctx context.Context, payload models.ConfirmTransactionRequest) (*models.ConfirmTransactionResponse, error) {
	logger.L().Info("Start confirm transaction",
		zap.Any("transaction id", payload.TransactionId),
	)

	transaction, err := s.transactionRepo.GetTransactionById(ctx, payload.TransactionId)
	if err != nil {
		logger.L().Error("failed to get transaction by id", zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appError.AppError{Code: http.StatusNotFound, Message: ErrTransactionNotFound}
		}
		return nil, appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// prevent confirm status
	if transaction.Status != entities.TransactionStatusEnumVerified {
		logger.L().Error("invalid transaction status", zap.Any("status", transaction.Status))
		return nil, appError.AppError{Code: http.StatusBadRequest, Message: ErrInvalidConfirmStatus}
	}

	now := time.Now().UTC()
	expired := transaction.ExpiresAt

	// prevent expired transaction
	if expired.Before(now) {
		logger.L().Error("transaction expired", zap.Any("expired_at", transaction.ExpiresAt))
		return nil, appError.AppError{Code: http.StatusBadRequest, Message: ErrTransactionExpred}
	}

	err = s.transactionRepo.ConfirmTransaction(ctx, *transaction)
	if err != nil {
		logger.L().Error("failed to update transaction", zap.Error(err))
		return nil, appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	newTransaction, err := s.transactionRepo.GetTransactionById(ctx, payload.TransactionId)
	if err != nil {
		logger.L().Error("failed to get transaction by id", zap.Error(err))
		return nil, appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &models.ConfirmTransactionResponse{
		TransactionId: newTransaction.TransactionID,
		UserId:        newTransaction.UserId,
		Amount:        newTransaction.Amount.InexactFloat64(),
		Status:        newTransaction.Status,
		Balance:       newTransaction.Wallet.Balance.InexactFloat64(),
	}, nil
}
