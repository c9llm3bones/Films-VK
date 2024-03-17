package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) InitRoutes() http.Handler {
	log.Println("Server has started")
	r := mux.NewRouter()

	r.HandleFunc("/auth", h.authentication).Methods("POST")
	return r
}
