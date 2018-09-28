package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dyong0/atm/internal/pkg/test"
	"github.com/gin-gonic/gin"
)

func TestAuthorize(t *testing.T) {
	router := gin.New()
	ConfigureAuthRouter(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/authorize", nil)
	router.ServeHTTP(w, req)
	res := test.ParseBody(w.Body)

	if expect, got := "token", res["access_token"]; expect != got {
		t.Errorf("Expect %v, got %v", expect, got)
	}
}
