package main

import (
	"errors"
	"strconv"

	"github.com/abiosoft/ishell"
	atmPkg "github.com/dyong0/atm/pkg/atm"
	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/dyong0/atm/pkg/atm/account/method"
	"github.com/dyong0/atm/pkg/atm/currency"
)

var (
	errUnknownAccountMethod = errors.New("unknown account method")
	errUnknownMenu          = errors.New("unknown menu")
	errATMInitiationFailure = errors.New("failed to initiate ATM")
	errNotSupportedMethod   = errors.New("not supported method")
	errNotSupportedAction   = errors.New("not supported action")
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
		return errUnknownAccountMethod
	}
}
func (da *dialogATM) readCard() error {
	da.dialog.Println("Insert card...")
	da.dialog.Println("Card No.: ")
	cardNo := da.dialog.ReadLine()
	da.dialog.Println("Password: ")
	password := da.dialog.ReadPassword()

	return da.atm.ReadAccount(method.NewCard(cardNo, password))
}
func (da *dialogATM) readAccountBank() error {
	return errNotSupportedMethod
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
	return errNotSupportedAction
}

func (da *dialogATM) printAccounts() {
	da.dialog.Println("=======Available accounts=======")
	for _, a := range da.accounts {
		da.dialog.Printf("ID: %5s PW: %5s\n", a.holderName, a.password)
	}
	da.dialog.Println("================================")
	da.dialog.Println("Welcome to ATM, type start")
}
