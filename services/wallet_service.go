package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/mahdidadashzadeh/wallet/models"
	"github.com/mahdidadashzadeh/wallet/repositories"
)

type WalletService struct {
	BalanceRepo     repositories.BalanceRepository
	TransactionRepo repositories.TransactionRepository
	mu              sync.Mutex
}

func (s *WalletService) GetBalance(userID int) (models.Balance, error) {
	return s.BalanceRepo.GetBalance(userID)
}

func (s *WalletService) Deposit(userID int, amount float64, description string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.BalanceRepo.GetBalance(userID)
	if err != nil {
		return err
	}
	balance.Balance += amount
	balance.UpdatedAt = time.Now()

	transaction := models.Transaction{
		UserID:         userID,
		CreatedAt:      time.Now(),
		Type:           models.Deposit,
		Status:         models.StatusSuccess,
		Amount:         amount,
		CurrentBalance: balance.Balance,
		Description:    description,
	}

	if err := s.BalanceRepo.UpdateBalance(balance); err != nil {
		return err
	}
	return s.TransactionRepo.CreateTransaction(transaction)
}

func (s *WalletService) Withdraw(userID int, amount float64, description string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.BalanceRepo.GetBalance(userID)
	if err != nil {
		return err
	}
	if balance.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	balance.Balance -= amount
	balance.UpdatedAt = time.Now()

	transaction := models.Transaction{
		UserID:         userID,
		CreatedAt:      time.Now(),
		Type:           models.Withdraw,
		Status:         models.StatusSuccess,
		Amount:         amount,
		CurrentBalance: balance.Balance,
		Description:    description,
	}

	if err := s.BalanceRepo.UpdateBalance(balance); err != nil {
		return err
	}
	return s.TransactionRepo.CreateTransaction(transaction)
}

func (s *WalletService) GetTransactions(userID int, page, limit int) ([]models.Transaction, error) {
	return s.TransactionRepo.GetTransactions(userID, page, limit)
}
