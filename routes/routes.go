package routes

import (
	"wallet-app/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, walletController *controllers.WalletController) {
	e.GET("/users/:id/balance", walletController.GetBalance)
	e.GET("/users/:id/transaction", walletController.GetTransactions)
	e.POST("/users/:id/deposit", walletController.Deposit)
	e.POST("/users/:id/withdraw", walletController.Withdraw)
}
