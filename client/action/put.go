package action

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/marc-sw/einkaufsliste/client/console"
	"github.com/marc-sw/einkaufsliste/core"
)

type PutAction struct {
}

func (action *PutAction) String() string {
	return "Edit product"
}

func (action *PutAction) Execute() {
	console.Info("Which product should be edited?")
	console.ListeAuf(Produkte)
	var index = console.LeseBestimmteZahl("Productnumber", 1, len(Produkte)) - 1
	var produkt core.Produkt = Produkte[index]
	console.Info("Enter new values or keep them empty to use the previos value.")
	produkt.Name = console.LeseOptionalString("Name", produkt.Name)
	produkt.Menge = console.LeseOptionalZahl("Quantity", produkt.Menge)
	produkt.Erledigt = console.LeseOptionalBool("Done", produkt.Erledigt)

	data, _ := json.Marshal(produkt)
	request, _ := http.NewRequest(http.MethodPut, Server+Endpoint, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)

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
	console.Info(string(body))
}
