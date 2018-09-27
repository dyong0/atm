package main

import (
	"errors"
	"strconv"

	"github.com/abiosoft/ishell"
	atmPkg "github.com/dyong0/atm/internal/pkg/atm"
	"github.com/dyong0/atm/internal/pkg/atm/account"
	"github.com/dyong0/atm/internal/pkg/atm/account/medium"
	"github.com/dyong0/atm/internal/pkg/atm/currency"
)

var (
	errUnknownAccountMedium = errors.New("unknown account medium")
	errUnknownMenu          = errors.New("unknown menu")
	errATMInitiationFailure = errors.New("Failed to initiate ATM")
	errNotSupportedMedium   = errors.New("not supported medium")
)

type accounts []struct {
	holderName string
	password   string
	balance    uint32
}

func newDialogATM(aa accounts, dialog ishell.Actions) (*dialogATM, error) {
	repo := account.NewMemRepository()
	for _, ia := range aa {
		a := account.NewAccount(currency.CurrencyKindYen)
		am, _ := currency.NewAmount(currency.CurrencyKindYen, ia.balance)
		a.Deposit(am)

		err := repo.Create(*a, ia.holderName, ia.password)
		if err != nil {
			dialog.Println(err.Error())
			return nil, errATMInitiationFailure
		}
	}

	atm, err := atmPkg.NewATM(repo)
	if err != nil {
		return nil, errATMInitiationFailure
	}

	return &dialogATM{
		atm:      atm,
		accounts: aa,
		dialog:   dialog,
	}, nil
}

type dialogATM struct {
	atm      *atmPkg.ATM
	dialog   ishell.Actions
	accounts accounts
}

func (da *dialogATM) createAccount(currencyKind currency.CurrencyKind, holderName string, password string) error {
	return nil
}
func (da *dialogATM) readAccount() error {
	choice := da.dialog.MultiChoice([]string{
		"Card",
		"Account Book",
	}, "Card or Account Book?")
	da.printAccounts()
	switch choice {
	case 0:
		return da.readCard()
	case 1:
		return da.readAccountBank()
	default:
		return errUnknownAccountMedium
	}
}
func (da *dialogATM) readCard() error {
	da.dialog.Println("Insert card...")
	da.dialog.Println("Card No.: ")
	cardNo := da.dialog.ReadLine()
	da.dialog.Println("Password: ")
	password := da.dialog.ReadPassword()

	return da.atm.ReadAccount(medium.NewCard(cardNo, password))
}
func (da *dialogATM) readAccountBank() error {
	return errNotSupportedMedium
}

func (da *dialogATM) processMenu() (bool, error) {
	choice := da.dialog.MultiChoice([]string{
		"Check balance",
		"Deposit",
		"Withdraw",
		"Exit",
	}, "Select menu")
	switch choice {
	case 0:
		da.balance()
		return true, nil
	case 1:
		return true, da.deposit()
	case 2:
		return true, da.withdraw()
	case 3:
		return false, nil
	default:
		return false, errUnknownMenu
	}
}
func (da *dialogATM) balance() {
	curKindName, _ := currency.CurrencyKindName(da.atm.CurrencyKind())
	da.dialog.Printf("Balance: %d %s\n", da.atm.Balance(), curKindName)
	da.dialog.Println("Enter to exit")
	da.dialog.ReadLine()
}
func (da *dialogATM) deposit() error {
	v, err := strconv.Atoi(da.dialog.ReadLine())
	if err != nil {
		da.dialog.Println("Invalid amount")
		da.dialog.Println("Enter to exit")
		da.dialog.ReadLine()
	}

	amt, err := currency.NewAmount(da.atm.CurrencyKind(), uint32(v))
	if err != nil {
		da.dialog.Println("Invalid amount")
		da.dialog.Println("Enter to exit")
		da.dialog.ReadLine()
	}

	err = da.atm.Deposit(amt)
	if err != nil {
		da.dialog.Println("Invalid amount")
		da.dialog.Println("Enter to exit")
		da.dialog.ReadLine()
	}

	return nil
}
func (da *dialogATM) withdraw() error {
	return nil
}

func (da *dialogATM) printAccounts() {
	da.dialog.Println("=======Available accounts=======")
	for _, a := range da.accounts {
		da.dialog.Printf("ID: %5s PW: %5s\n", a.holderName, a.password)
	}
	da.dialog.Println("================================")
	da.dialog.Println("Welcome to ATM, type start")
}
