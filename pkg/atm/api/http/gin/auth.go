package gin

import (
	"errors"
	"net/http"
	"time"

	"github.com/dyong0/atm/pkg/atm/account"

	ares "github.com/dyong0/atm/pkg/atm/api/http/response"
	"github.com/dyong0/atm/pkg/atm/auth/token"
	"github.com/gin-gonic/gin"
)

type authRouterDependencies struct {
	accRepo   account.Repository
	tokenRepo token.Repository
}

func AuthRouterDependencies() authRouterDependencies {
	return authRouterDependencies{}
}

func (deps authRouterDependencies) WithAccountRepository(r account.Repository) authRouterDependencies {
	deps.accRepo = r
	return deps
}

func (deps authRouterDependencies) WithTokenRepository(r token.Repository) authRouterDependencies {
	deps.tokenRepo = r
	return deps
}

func ConfigureAuthRouter(r gin.IRouter, deps authRouterDependencies) {
	r.POST("/authorize", authorize(deps.accRepo, deps.tokenRepo))
}

func authorize(accRepo account.Repository, tokenRepo token.Repository) gin.HandlerFunc {
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
			c.JSON(ares.BadRequest())
			return
		}

		acc, err := accRepo.Account(req.AccountID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ares.Error().Withmessage("invalid credential"))
			return
		}
		err = acc.Authenticate(req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ares.Error().Withmessage("invalid credential"))
			return
		}

		newToken, err := tokenRepo.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage("failed to issue a token"))
			return
		}
		newToken.Authorization["account"] = acc
		err = tokenRepo.Save(newToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ares.Error().Withmessage("failed to issue a token"))
			return
		}

		c.JSON(http.StatusOK, response{
			AccessToken: newToken.ID,
			ExpiresAt:   newToken.ExpiresAt.Format(time.UnixDate),
		})
	}
}

func mustWithValidToken(tokenRepo token.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := tokenRepo.Token(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		}

		c.Keys["token"] = tok

		c.Next()
	}
}

func withAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tok token.Token
		tok, ok := c.Keys["token"].(token.Token)
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
			return
		}
		acc, ok := tok.Authorization["account"]
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
			return
		}

		c.Keys["account"] = acc

		c.Next()
	}
}
