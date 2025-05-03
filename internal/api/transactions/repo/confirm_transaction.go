package repo

import (
	"api-wallet/internal/entities"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *TransactionRepository) ConfirmTransaction(ctx context.Context, transaction entities.Transaction) error {
	return r.postgres.Transaction(func(tx *gorm.DB) error {
		var transactionResp entities.Transaction

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("transaction_id = ?", transaction.TransactionID).First(&transactionResp).Error; err != nil {
			return err
		}

		var wallet entities.Wallet
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("wallet_id = ?", transaction.WalletId).First(&wallet).Error; err != nil {
			return err
		}

		if wallet.Balance.LessThan(transaction.Amount) {
			return errors.New("insufficient balance")
		}

		if transactionResp.Status != entities.TransactionStatusEnumVerified {
			return errors.New("transaction is already update")
		}

		// Update transaction status
		if err := tx.Model(&entities.Transaction{}).
			Where("transaction_id = ?", transactionResp.TransactionID).
			Update("status", entities.TransactionStatusEnumCompleted).Error; err != nil {
			return err
		}

		// Deduct the balance
		newBalance := wallet.Balance.Sub(transaction.Amount)
		if err := tx.Model(&entities.Wallet{}).
			Where("wallet_id = ?", wallet.WalletId).
			Update("balance", newBalance).Error; err != nil {
			return err
		}

		return nil
	})
}
