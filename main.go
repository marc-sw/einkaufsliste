package main

import (
	"log"
	"os"

	"github.com/marc-sw/einkaufsliste/api"
	"github.com/marc-sw/einkaufsliste/client"
)

// getbyid erstmal nicht, da kein sinn und keine prio
// todo: execute methode von post, put und delete

func main() {
	if len(os.Args) > 1 && os.Args[1] == "client" {
		log.Println("client wird gestartet")
		client.Run()
	} else {
		log.Println("rest-server wird gestartet")
		api.Run()
	}
}
