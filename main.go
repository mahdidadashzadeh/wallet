package main

import (
	"wallet-app/config"
	"wallet-app/controllers"
	"wallet-app/routes"
	"wallet-app/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := config.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}

	walletService := &services.WalletService{DB: db}
	walletController := &controllers.WalletController{Service: walletService}

	routes.SetupRoutes(e, walletController)

	e.Start(":8080")
}
