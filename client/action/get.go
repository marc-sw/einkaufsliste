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
	return "Get all products"
}

func (action *GetAction) Execute() {
	response, err := http.DefaultClient.Get(Server + Endpoint)
	if err != nil {
		console.Error("No connection to server possible.")
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
		console.Error("Could not read response from server")
		return
	}
	Produkte = produkte
	console.Info("Shoppinglist")
	console.ListeAuf(produkte)
}
