package action

import (
	"fmt"
	"log"
	"os"

	"github.com/marc-sw/einkaufsliste/core"
)

var (
	Server   string = "http://localhost:8090"
	Endpoint string = "/produkt"
	Produkte []core.Produkt
)

type Action interface {
	fmt.Stringer
	Execute()
}

type BeendenAction struct {
}

func (action *BeendenAction) String() string {
	return "Client beenden"
}

func (action *BeendenAction) Execute() {
	log.Println("Client wurde beendet")
	os.Exit(0)
}

func CreateAllActions() []Action {
	return []Action{
		&GetAction{},
		&PostAction{},
		&PutAction{},
		&DeleteAction{},
		&BeendenAction{},
	}
}
