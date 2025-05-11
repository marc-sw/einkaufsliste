package action

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marc-sw/einkaufsliste/client/console"
	"github.com/marc-sw/einkaufsliste/core"
)

type GetAction struct {
}

func (action *GetAction) String() string {
	return "Alle Produkte abfragen"
}

func (action *GetAction) Execute() {
	response, err := http.DefaultClient.Get(Server + Endpoint)
	if err != nil {
		console.Error("Anfrage konnte nicht an den Server gestellt werden.")
		return
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		console.Error(string(body))
		return
	}

	var produkte []core.Produkt
	if err := json.Unmarshal(body, &produkte); err != nil {
		console.Error("Antwort vom Server konnte nicht gelesen werden.")
		return
	}
	Produkte = produkte
	console.Info("Einkaufsliste")
	console.ListeAuf(produkte)
}
