package models

type Balance struct {
	UserID   int     `gorm:"primaryKey;not null" json:"user_id"`
	Balance  float64 `gorm:"not null" json:"balance"`
	UpdateAt string  `gorm:"not null" json:"update_at"`
}

type TransactionHistory struct {
	HistoryID      int     `gorm:"primaryKey;not null" json:"history_id"`
	UserID         int     `gorm:"not null" json:"user_id"`
	CreateAt       string  `gorm:"not null" json:"create_at"`
	Type           string  `gorm:"type:enum('deposit', 'withdraw');not null" json:"type"` // "deposit" or "withdraw"
	Status         int     `gorm:"not null" json:"status"`
	Amount         float64 `gorm:"not null" json:"amount"`
	CurrentBalance float64 `gorm:"not null" json:"current_balance"`
	Description    string  `gorm:"type:varchar(45)" json:"description"`
}
