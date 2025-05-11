package main

import (
	"log"
	"os"

	"github.com/marc-sw/einkaufsliste/api"
	"github.com/marc-sw/einkaufsliste/client"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "client" {
		log.Println("client wird gestartet")
		client.Run()
	} else {
		log.Println("rest-server wird gestartet")
		api.Run()
	}
}
