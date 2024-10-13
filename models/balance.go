package models

import "time"

type Balance struct {
	UserID    int       `gorm:"primaryKey;not null"`
	Balance   float64   `gorm:"type:decimal(5,2);not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
