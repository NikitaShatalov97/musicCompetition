package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NikitaShatalov97/musicCompetition.git/internal/model"
	"github.com/go-chi/chi/v5"
)

type HelloHandler struct {
}

func CreateHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	resp := model.HelloResponce{
		Message: fmt.Sprintf("Hello, %s", name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
