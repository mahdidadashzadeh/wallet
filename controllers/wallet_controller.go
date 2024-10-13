package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mahdidadashzadeh/wallet/services"
)

type WalletController struct {
	Service services.WalletService
}

func (c *WalletController) GetBalance(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	balance, err := c.Service.GetBalance(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, balance)
}

func (c *WalletController) FetchTransactions(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	transactions, err := c.Service.GetTransactions(userID, page, limit)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, transactions)
}

func (c *WalletController) Deposit(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	var req struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err := c.Service.Deposit(userID, req.Amount, req.Description)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Deposit successful")
}

func (c *WalletController) Withdraw(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	var req struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err := c.Service.Withdraw(userID, req.Amount, req.Description)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Withdrawal successful")
}
