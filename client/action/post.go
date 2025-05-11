package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marc-sw/einkaufsliste/client/console"
	"github.com/marc-sw/einkaufsliste/core"
)

type PostAction struct {
}

func (action *PostAction) String() string {
	return "Produkt hinzufuegen"
}

func (action *PostAction) Execute() {
	var produkt core.Produkt
	fmt.Println("Neues Produkt hinzufuegen")
	produkt.Name = console.LeseString("Name")
	produkt.Menge = console.LeseZahl("Menge")

	data, _ := json.Marshal(produkt)

	response, err := http.DefaultClient.Post(Server+Endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		console.Error("Anfrage konnte nicht an den Server gestellt werden.")
		return
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusCreated {
		console.Error(string(body))
		return
	}
	console.Info(string(body))
}
