package gin

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/dyong0/atm/pkg/atm/auth/token"
	"github.com/dyong0/atm/pkg/atm/auth/token/generator"
	"github.com/dyong0/atm/pkg/atm/currency"

	"github.com/dyong0/atm/internal/pkg/test"
	"github.com/gin-gonic/gin"
)

func TestAuthorize(t *testing.T) {
	router := gin.New()
	accRepo := account.NewMemRepository()
	accRepo.Create(*account.NewAccount(currency.CurrencyKindYen), "id", "pw")
	genMock := genMock{}

	payload := `
    {
      "account_id": "id",
      "password": "pw"
	}`
	router.POST("/authorize", authorize(accRepo, token.NewMemRepository(genMock)))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/authorize", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	now := time.Now()
	router.ServeHTTP(w, req)
	res := test.ParseBody(w.Body)

	if expect, got := "token", res["access_token"]; expect != got {
		t.Errorf("Expect %s, got %s", expect, got)
	}

	if expect, got := now.Add(time.Minute*5).Format(time.UnixDate), res["expires_at"]; expect != got {
		t.Errorf("Expect token expires at %s, got %s", expect, got)
	}
}

type genMock struct {
	generator.Generator
}

func (g genMock) Generate() (string, error) {
	return "token", nil
}
