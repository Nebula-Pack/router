package main

import (
	"log"
	"net/http"

	"github.com/Nebula-Pack/router/pkg/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Server starting on :2321")
	log.Fatal(http.ListenAndServe(":2321", r))
}
