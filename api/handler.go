package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/marc-sw/einkaufsliste/core"
)

var repository = core.MockProductRepository()
var ungueltigeIdError = errors.New("Unbekannte Produkt Id")

func handleInternalError(err error, w http.ResponseWriter) {
	http.Error(w, "Ein unerwarteter Fehler ist aufgetreten", http.StatusInternalServerError)
	log.Println("Ein interner Fehler ist aufgetreten: " + err.Error())
}

func handleClientError(err error, w http.ResponseWriter) {
	http.Error(w, "Ungueltige Eingabe: ", http.StatusBadRequest)
	log.Panicln("Ungueltige Eingabe vom Client: " + err.Error())
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	produkte := repository.GetAllProdukte()
	data, err := json.Marshal(produkte)

	if err != nil {
		handleInternalError(err, w)
	} else {
		w.Write(data)
		log.Println("Client hat alle Produkte abgefragt")
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var produkt core.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		handleClientError(err, w)
	} else {
		repository.CreateProdukt(produkt)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Das Produkt wurde der Einkaufsliste hinzugefuegt."))
		log.Println("Client hat ein Produkt der Einkaufsliste hinzugefuegt.")
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var produkt core.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		handleClientError(err, w)
	} else {
		var bearbeitet = repository.UpdateProdukt(produkt)

		if bearbeitet {
			w.Write([]byte("Das Produkt wurde bearbeitet."))
			log.Println("Client hat ein Produkt bearbeitet.")
		} else {
			handleClientError(ungueltigeIdError, w)
		}
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := parseProduktId(r)
	var geloescht = false
	if id != -1 {
		geloescht = repository.DeleteProdukt(id)
	}

	if geloescht {
		log.Println("Das Produkt wurde aus der Einkaufsliste entfernt.")
		w.Write([]byte("Client hat ein Produkt aus der Einkaufsliste entfernt."))
	} else {
		handleClientError(ungueltigeIdError, w)
	}

}
