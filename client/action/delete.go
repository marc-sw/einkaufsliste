package action

import (
	"fmt"
	"io"
	"net/http"

	"github.com/marc-sw/einkaufsliste/client/console"
	"github.com/marc-sw/einkaufsliste/core"
)

type DeleteAction struct {
}

func (action *DeleteAction) String() string {
	return "Remove product"
}

func (action *DeleteAction) Execute() {
	console.Info("Which product should be removed?")
	console.ListeAuf(Produkte)
	var index = console.LeseBestimmteZahl("Productnumber", 1, len(Produkte)) - 1
	var produkt core.Produkt = Produkte[index]
	var url = fmt.Sprintf("%s/%d", Server+Endpoint, produkt.Id)
	request, _ := http.NewRequest(http.MethodDelete, url, nil)

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
