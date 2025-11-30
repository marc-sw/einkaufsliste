package api

import (
	"net/http"
	"strconv"
)

func Run() {
	fileServer := http.FileServer(http.Dir("./svelte/client/dist"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /produkt", handleGetAll)
	mux.HandleFunc("POST /produkt", handlePost)
	mux.HandleFunc("PUT /produkt", handlePut)
	mux.HandleFunc("DELETE /produkt/{id}", handleDelete)
	mux.Handle("/", fileServer)

	http.ListenAndServe(":8090", mux)
}

func parseProduktId(r *http.Request) int {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return -1
	}
	return id
}
