package gin

import (
	"errors"
	"net/http"

	"github.com/dyong0/atm/pkg/atm"
	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/gin-gonic/gin"
)

func ConfigureAccountRouter(r gin.IRouter, atm *atm.ATM) {
	r.Use(withAccount())
	r.GET("/balance", balance(atm))
	r.POST("/deposit", deposit(atm))
	r.POST("/withdraw", withdraw(atm))
}

func withAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, err := accountByToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
			return
		}

		c.Set("account", a)

		c.Next()
	}
}

func accountByToken(authHeader string) (account.Account, error) {
	return account.Account{}, nil
}
