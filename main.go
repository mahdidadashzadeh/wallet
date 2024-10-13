package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mahdidadashzadeh/wallet/config"
	"github.com/mahdidadashzadeh/wallet/controllers"
	"github.com/mahdidadashzadeh/wallet/repositories"
	"github.com/mahdidadashzadeh/wallet/services"
)

func main() {
	config.ConnectDatabase()

	e := echo.New()

	balanceRepo := repositories.BalanceRepository{DB: config.DB}
	transactionRepo := repositories.TransactionRepository{DB: config.DB}
	walletService := services.WalletService{
		BalanceRepo:     balanceRepo,
		TransactionRepo: transactionRepo,
	}
	walletController := controllers.WalletController{Service: walletService}

	e.GET("/users/:id/balance", walletController.GetBalance)
	e.GET("/users/:id/transactions", walletController.FetchTransactions)
	e.POST("/users/:id/deposit", walletController.Deposit)
	e.POST("/users/:id/withdraw", walletController.Withdraw)

	e.Start(":8080")
}
