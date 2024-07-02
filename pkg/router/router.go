package router

import (
	"github.com/Nebula-Pack/router/internal/handler"
	"github.com/gorilla/mux"
)

// NewRouter creates a new router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/{key}", handler.KeyHandler).Methods("GET")
	return r
}
