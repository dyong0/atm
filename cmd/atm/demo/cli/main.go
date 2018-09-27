package main

import (
	"github.com/abiosoft/ishell"
)

var atm *dialogATM

func main() {
	shell := ishell.New()

	var err error
	atm, err = newDialogATM(accounts{
		{"me", "pw", 719180},
		{"poor", "man", 100},
		{"i", "donno", 6551621},
	}, shell)
	if err != nil {
		shell.Println(err.Error())
		return
	}

	shell.Println("Type start to begin")

	shell.AddCmd(&ishell.Cmd{
		Name: "start",
		Help: "Start",
		Func: handleStart,
	})

	shell.Run()
}

func handleStart(ctx *ishell.Context) {
	ctx.ShowPrompt(false)
	defer ctx.ShowPrompt(true)
	defer func() {
		ctx.Println("Type start to begin another")
	}()

	err := atm.readAccount()
	if err != nil {
		ctx.Println(err.Error())
		return
	}

	for {
		processNext, err := atm.processMenu()
		if err != nil {
			ctx.Println(err.Error())
			ctx.Println("Enter to exit")
			ctx.ReadLine()
		}
		if !processNext {
			break
		}
	}
}
