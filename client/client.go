package client

import (
	"github.com/marc-sw/einkaufsliste/client/action"
	"github.com/marc-sw/einkaufsliste/client/console"
)

func Run() {
	var actions []action.Action = action.CreateAllActions()
	var actionId int

	for {
		console.PrintSeperatorLine()
		console.ListeAuf(actions)
		actionId = console.LeseBestimmteZahl("Aktionsnummer", 1, len(actions)) - 1 //leseNutzerEingabe(1, len(actions)) - 1
		console.PrintSeperatorLine()
		actions[actionId].Execute()
	}
}
