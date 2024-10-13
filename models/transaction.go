package models

import "time"

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type TransactionStatus int

const (
	StatusSuccess TransactionStatus = 1
	StatusFailed  TransactionStatus = 2
)

type Transaction struct {
	HistoryID      int               `gorm:"primaryKey;not null"`
	UserID         int               `gorm:"not null"`
	CreatedAt      time.Time         `gorm:"not null"`
	Type           TransactionType   `gorm:"type:ENUM('deposit', 'withdraw');not null"`
	Status         TransactionStatus `gorm:"type:ENUM(1, 2);not null"`
	Amount         float64           `gorm:"type:decimal(5,2);not null"`
	CurrentBalance float64           `gorm:"not null"`
	Description    string            `gorm:"type:varchar(45)"`
}
