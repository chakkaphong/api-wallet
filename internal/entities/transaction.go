package entities

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionStatusEnum string
type PaymentMethodEnum string

const (
	// Status
	TransactionStatusEnumVerified  TransactionStatusEnum = "verified"
	TransactionStatusEnumCompleted TransactionStatusEnum = "completed"

	// Payment Methods
	PaymentMethodEnumCreditCard PaymentMethodEnum = "credit_card"
)

type Transaction struct {
	TransactionID string
	WalletId      int
	UserId        int
	PaymentMethod PaymentMethodEnum
	Status        TransactionStatusEnum
	Amount        decimal.Decimal
	CreatedAt     time.Time
	ExpiresAt     time.Time

	//FK
	Wallet *Wallet `gorm:"foreignKey:WalletId;references:WalletId"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	if t.TransactionID == "" {
		t.TransactionID = fmt.Sprintf("tx%d", time.Now().UnixNano())
	}

	return nil
}
