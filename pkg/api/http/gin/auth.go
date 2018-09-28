package gin

import (
	"net/http"
	"time"

	ahttp "github.com/dyong0/atm/pkg/api/http"
	"github.com/dyong0/atm/pkg/atm"
	"github.com/dyong0/atm/pkg/atm/account/method"
	"github.com/dyong0/atm/pkg/atm/auth/token"
	"github.com/gin-gonic/gin"
)

func ConfigureAuthRouter(r gin.IRouter, atm *atm.ATM, tokenRepo token.Repository) {
	r.POST("/authorize", authorize(atm, tokenRepo))
}

func authorize(atm *atm.ATM, tokenRepo token.Repository) gin.HandlerFunc {
	type request struct {
		AccountID string `json:"account_id" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}
	type response struct {
		AccessToken string `json:"access_token"`
		ExpiresAt   string `json:"expires_at"`
	}

	return func(c *gin.Context) {
		var req request
		err := c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ahttp.ErrorResponse().Withmessage("invalid credential"))
			return
		}

		accountMethod := method.NewPlain(req.AccountID, req.Password)
		err = atm.ReadAccount(accountMethod)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ahttp.ErrorResponse().Withmessage("invalid credential"))
			return
		}

		newToken, err := tokenRepo.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ahttp.ErrorResponse().Withmessage("failed to issue a token"))
			return
		}
		newToken.Authorization = accountMethod
		err := tokenRepo.Save(newToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ahttp.ErrorResponse().Withmessage("failed to issue a token"))
			return
		}

		c.JSON(http.StatusOK, response{
			AccessToken: newToken.String(),
			ExpiresAt:   time.Now().String(),
		})
	}
}
