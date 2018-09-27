package gin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ConfigureAuthRouter(r gin.IRouter) {
	r.POST("/authorize", authorize())
}

func authorize() func(c *gin.Context) {
	type request struct {
	}
	type response struct {
		AccessToken string `json:"access_token"`
		ExpiresAt   string `json:"expires_at"`
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, response{
			AccessToken: "token",
			ExpiresAt:   time.Now().String(),
		})
	}
}
