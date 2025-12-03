package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/marc-sw/einkaufsliste/core"
)

var repository = core.MockProductRepository()
var ungueltigeIdError = errors.New("Unknown product id")

func handleInternalError(err error, w http.ResponseWriter) {
	http.Error(w, "An unknown error happened", http.StatusInternalServerError)
	log.Println("An internal error happened: " + err.Error())
}

func handleClientError(err error, w http.ResponseWriter) {
	http.Error(w, "Invalid input: ", http.StatusBadRequest)
	log.Panicln("Invalid input from client: " + err.Error())
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	produkte := repository.GetAllProdukte()
	data, err := json.Marshal(produkte)

	if err != nil {
		handleInternalError(err, w)
	} else {
		w.Write(data)
		log.Println("Client fetched all products")
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var produkt core.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		handleClientError(err, w)
	} else {
		repository.CreateProdukt(produkt)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product added to shopping list."))
		log.Println("Client added a product to the shopping list.")
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var produkt core.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		handleClientError(err, w)
	} else {
		var bearbeitet = repository.UpdateProdukt(produkt)

		if bearbeitet {
			w.Write([]byte("Product edited."))
			log.Println("Client edited a product.")
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
		log.Println("Product removed from shopping list.")
		w.Write([]byte("Client removed a product from the shopping list."))
	} else {
		handleClientError(ungueltigeIdError, w)
	}

}
