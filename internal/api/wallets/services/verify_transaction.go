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

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ErrUserNotFound        string = "user not found"
	ErrWalletNotFound      string = "wallet not found"
	ErrInsufficientBalance string = "insufficient balance"
)

func (s *WalletServices) VerifyTransaction(ctx context.Context, payload models.VerifyRequest) (*entities.Transaction, error) {
	logger.L().Info("Start verifying payment",
		zap.Any("method", payload.PaymentMethod),
		zap.Any("amount", payload.Amount),
	)

	user, err := s.userRepo.GetUserByID(ctx, payload.UserId)
	if err != nil {
		logger.L().Error("failed to get user by id", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appError.AppError{Code: http.StatusNotFound, Message: ErrUserNotFound}
		}
		return nil, err
	}

	cusWallet, err := s.repo.GetWalletByUserId(ctx, user.UserId)
	if err != nil {
		logger.L().Error("failed to get wallet by id", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appError.AppError{Code: http.StatusNotFound, Message: ErrWalletNotFound}
		}
		return nil, err
	}

	now := time.Now().UTC()
	amountDecimal := decimal.NewFromFloat(payload.Amount)

	if cusWallet.IsInsufficientBalance(amountDecimal) {
		logger.L().Error("insufficient balance", zap.Any("balance", cusWallet.Balance))
		return nil, appError.AppError{Code: http.StatusBadRequest, Message: ErrInsufficientBalance}
	}

	txPayload := entities.Transaction{
		WalletId:      cusWallet.WalletId,
		UserId:        payload.UserId,
		PaymentMethod: entities.PaymentMethodEnumCreditCard,
		Status:        entities.TransactionStatusEnumVerified,
		Amount:        amountDecimal,
		ExpiresAt:     now.Add(1 * time.Minute),
	}
	transactionCreated, err := s.transactionRepo.SaveTransaction(ctx, txPayload)
	if err != nil {
		logger.L().Error("failed to create transaction", zap.Error(err))
		return nil, err
	}

	return transactionCreated, nil
}
