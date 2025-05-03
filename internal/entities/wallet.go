package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	WalletId  int
	UserId    int
	Balance   decimal.Decimal
	CreatedAt time.Time

	// FK
	User *User `gorm:"foreignKey:UserId;references:UserId"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) IsInsufficientBalance(amount decimal.Decimal) bool {
	return w.Balance.LessThan(amount)
}
