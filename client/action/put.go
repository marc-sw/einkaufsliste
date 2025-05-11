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
	return "Produkt bearbeiten"
}

func (action *PutAction) Execute() {
	console.Info("Welches Produkt soll bearbeitet werden?")
	console.ListeAuf(Produkte)
	var index = console.LeseBestimmteZahl("Produktnummer", 1, len(Produkte)) - 1
	var produkt core.Produkt = Produkte[index]
	console.Info("Gebe neue Werte ein oder lasse die Eingabe weg.")
	produkt.Name = console.LeseOptionalString("Name", produkt.Name)
	produkt.Menge = console.LeseOptionalZahl("Menge", produkt.Menge)
	produkt.Erledigt = console.LeseOptionalBool("Erledigt", produkt.Erledigt)

	data, _ := json.Marshal(produkt)
	request, _ := http.NewRequest(http.MethodPut, Server+Endpoint, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)

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
	console.Info(string(body))
}
