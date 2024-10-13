package repositories

import (
	"github.com/mahdidadashzadeh/wallet/models"
	"gorm.io/gorm"
)

type BalanceRepository struct {
	DB *gorm.DB
}

func (r *BalanceRepository) GetBalance(userID int) (models.Balance, error) {
	var balance models.Balance
	err := r.DB.First(&balance, userID).Error
	return balance, err
}

func (r *BalanceRepository) UpdateBalance(balance models.Balance) error {
	return r.DB.Save(&balance).Error
}
