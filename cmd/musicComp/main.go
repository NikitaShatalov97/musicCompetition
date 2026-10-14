package main

import (
	"net/http"

	"github.com/NikitaShatalov97/musicCompetition.git/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	handleHello := handler.CreateHelloHandler()
	r := chi.NewRouter()

	r.Get("/{name}", handleHello.Hello)

	http.ListenAndServe(":8080", r)

}
