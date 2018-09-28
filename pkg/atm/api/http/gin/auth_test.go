package gin

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	genMock := genMock{}
	deps := AuthRouterDependencies()
	deps = deps.WithAccountRepository(accRepo)
	deps = deps.WithTokenRepository(token.NewMemRepository(genMock))
	accRepo.Create(*account.NewAccount(currency.CurrencyKindYen), "id", "pw")
	ConfigureAuthRouter(router, deps)
	w := httptest.NewRecorder()
	payload := `
    {
      "account_id": "id",
      "password": "pw"
	}`
	req, _ := http.NewRequest("POST", "/authorize", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	res := test.ParseBody(w.Body)

	if expect, got := "token", res["access_token"]; expect != got {
		t.Errorf("Expect %v, got %v", expect, got)
	}
}

type genMock struct {
	generator.Generator
}

func (g genMock) Generate() (string, error) {
	return "token", nil
}
