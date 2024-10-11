package services

import (
	"errors"
	"time"
	"wallet-app/models"

	"gorm.io/gorm"
)

type WalletService struct {
	DB *gorm.DB
}

func (ws *WalletService) GetBalance(userID int) (models.Balance, error) {
	var balance models.Balance
	result := ws.DB.First(&balance, userID)
	return balance, result.Error
}

func (ws *WalletService) GetTransactions(userID int) ([]models.TransactionHistory, error) {
	var transactions []models.TransactionHistory
	result := ws.DB.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}

func (ws *WalletService) Deposit(userID int, amount float64, description string) error {
	var balance models.Balance
	result := ws.DB.First(&balance, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			balance = models.Balance{
				UserID:   userID,
				Balance:  amount, // موجودی اولیه برابر با مبلغ واریزی
				UpdateAt: time.Now().Format(time.RFC3339),
			}
			ws.DB.Create(&balance)
		} else {
			return result.Error
		}
	} else {
		balance.Balance += amount
		balance.UpdateAt = time.Now().Format(time.RFC3339)
		ws.DB.Save(&balance)
	}

	// ثبت تراکنش
	transaction := models.TransactionHistory{
		UserID:         userID,
		CreateAt:       time.Now().Format(time.RFC3339),
		Type:           "deposit",
		Status:         1,
		Amount:         amount,
		CurrentBalance: balance.Balance,
		Description:    description,
	}
	ws.DB.Create(&transaction)

	return nil
}

func (ws *WalletService) Withdraw(userID int, amount float64, description string) error {
	var balance models.Balance
	result := ws.DB.First(&balance, userID)
	if result.Error != nil {
		return result.Error
	}

	if balance.Balance < amount {
		return errors.New("insufficient funds")
	}

	// به‌روز کردن موجودی
	balance.Balance -= amount
	balance.UpdateAt = time.Now().Format(time.RFC3339)
	ws.DB.Save(&balance)

	// ثبت تراکنش
	transaction := models.TransactionHistory{
		UserID:         userID,
		CreateAt:       time.Now().Format(time.RFC3339),
		Type:           "withdraw",
		Status:         1,
		Amount:         amount,
		CurrentBalance: balance.Balance,
		Description:    description,
	}
	ws.DB.Create(&transaction)

	return nil
}
