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
	return "Add product"
}

func (action *PostAction) Execute() {
	var produkt core.Produkt
	fmt.Println("Add new product")
	produkt.Name = console.LeseString("Name")
	produkt.Menge = console.LeseZahl("Quantity")

	data, _ := json.Marshal(produkt)

	response, err := http.DefaultClient.Post(Server+Endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		console.Error("No connection to server possible.")
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
