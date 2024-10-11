package controllers

import (
	"net/http"
	"strconv"
	"wallet-app/services"

	"github.com/labstack/echo/v4"
)

type WalletController struct {
	Service *services.WalletService
}

func (wc *WalletController) GetBalance(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	balance, err := wc.Service.GetBalance(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, balance)
}

func (wc *WalletController) GetTransactions(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	transactions, err := wc.Service.GetTransactions(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactions)
}

func (wc *WalletController) Deposit(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var req struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := wc.Service.Deposit(userID, req.Amount, req.Description); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deposit successful")
}

func (wc *WalletController) Withdraw(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var req struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := wc.Service.Withdraw(userID, req.Amount, req.Description); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Withdrawal successful")
}
