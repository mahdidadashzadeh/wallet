package repositories

import (
	"github.com/mahdidadashzadeh/wallet/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r *TransactionRepository) CreateTransaction(transaction models.Transaction) error {
	return r.DB.Create(&transaction).Error
}

func (r *TransactionRepository) GetTransactions(userID int, page, limit int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where("user_id = ?", userID).Limit(limit).Offset((page - 1) * limit).Find(&transactions).Error
	return transactions, err
}
