package atm

import (
	"testing"

	"github.com/dyong0/atm/pkg/account"
)

func TestATM(t *testing.T) {
	atm := New()

	mockAccService := &MockAccountService{
		ReadAccountFunc: func(id string, pw string) error {
			return account.ErrAccountNotFound
		},
	}

	atm.accountService = mockAccService

	err := atm.ReadAccount("non existing", "pw")

	if err != account.ErrAccountNotFound {
		t.Error("Expect account to be not found")
	}
}

type MockAccountService struct {
	account.Service
	ReadAccountFunc func(string, string) error
}

func (s *MockAccountService) ReadAccount(id string, pw string) error {
	return s.ReadAccountFunc(id, pw)
}
