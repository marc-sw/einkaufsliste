package api

import (
	"net/http"
	"strconv"
)

func Run() {
	fileServer := http.FileServer(http.Dir("./svelte/client/dist"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /shoppinglist-api/v1/products", handleGetAll)
	mux.HandleFunc("POST /shoppinglist-api/v1/products", handlePost)
	mux.HandleFunc("PUT /shoppinglist-api/v1/products", handlePut)
	mux.HandleFunc("DELETE /shoppinglist-api/v1/products/{id}", handleDelete)
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
