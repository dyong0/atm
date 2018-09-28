package gin

import (
	"fmt"
	"net/http"

	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/dyong0/atm/pkg/atm/auth/token"
	"github.com/dyong0/atm/pkg/atm/currency"
	"github.com/gin-gonic/gin"

	ares "github.com/dyong0/atm/pkg/atm/api/http/response"
)

func ConfigureAccountRouter(r gin.IRouter, tokenRepo token.Repository, accRepo account.Repository) {
	r.Use(mustWithValidToken(tokenRepo), withAccount())
	r.GET("/balance", balance())
	r.POST("/deposit", deposit(accRepo))
	r.POST("/withdraw", withdraw(accRepo))
}

func balance() gin.HandlerFunc {
	type response struct {
		Balance      uint32 `json:"balance"`
		CurrencyKind string `json:"currency_kind"`
	}

	return func(c *gin.Context) {
		acc := accountFromContext(c)
		curkindName, err := currency.CurrencyKindName(acc.CurrencyKind())
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage("invalid currency kind"))
		}

		c.JSON(http.StatusOK, response{
			Balance:      acc.Balance(),
			CurrencyKind: curkindName,
		})
	}
}

func deposit(accRepo account.Repository) gin.HandlerFunc {
	type request struct {
		Amount       uint32 `json:"amount" binding:"required"`
		CurrencyKind string `json:"currency_kind" binding:"required"`
	}
	type response struct {
		AmountDeposited uint32 `json:"amount_deposited"`
		CurrencyKind    string `json:"currency_kind"`
	}

	return func(c *gin.Context) {
		var req request
		err := c.Bind(&req)
		if err != nil {
			c.JSON(ares.BadRequest())
			return
		}

		acc := accountFromContext(c)
		curkindName, err := currency.CurrencyKindName(acc.CurrencyKind())
		if err != nil {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid currency kind"))
		}
		if curkindName != req.CurrencyKind {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid currency kind"))
		}

		amt, err := currency.NewAmount(acc.CurrencyKind(), req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid amount"))
		}
		err = acc.Deposit(amt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage(fmt.Sprintf("failed to deposit amount %d %s", req.Amount, req.CurrencyKind)))
		}

		err = accRepo.Update(acc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage(fmt.Sprintf("failed to deposit amount %d %s", req.Amount, req.CurrencyKind)))
		}

		c.JSON(http.StatusOK, response{
			AmountDeposited: amt.Total(),
			CurrencyKind:    curkindName,
		})
	}
}

func withdraw(accRepo account.Repository) gin.HandlerFunc {
	type request struct {
		Amount       uint32 `json:"amount" binding:"required"`
		CurrencyKind string `json:"currency_kind" binding:"required"`
	}
	type response struct {
		WithdrawnAmount uint32 `json:"withdrawn_amount"`
		CurrencyKind    string `json:"currency_kind"`
	}

	return func(c *gin.Context) {
		var req request
		err := c.Bind(&req)
		if err != nil {
			c.JSON(ares.BadRequest())
			return
		}

		acc := accountFromContext(c)
		curkindName, err := currency.CurrencyKindName(acc.CurrencyKind())
		if err != nil {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid currency kind"))
		}
		if curkindName != req.CurrencyKind {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid currency kind"))
		}

		amt, err := currency.NewAmount(acc.CurrencyKind(), req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, ares.Error().Withmessage("invalid amount"))
		}
		withdrawn, err := acc.Withdraw(amt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage(fmt.Sprintf("failed to withdraw amount %d %s", req.Amount, req.CurrencyKind)))
		}

		err = accRepo.Update(acc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage(fmt.Sprintf("failed to withdraw amount %d %s", req.Amount, req.CurrencyKind)))
		}

		c.JSON(http.StatusOK, response{
			WithdrawnAmount: withdrawn.Total(),
			CurrencyKind:    curkindName,
		})
	}
}

func accountFromContext(c *gin.Context) account.Account {
	return c.Keys["account"].(account.Account)
}
