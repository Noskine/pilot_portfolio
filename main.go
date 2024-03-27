package main

import (
	"log"
	"net/http"

	"github.com/noskine/pilot_api/config"
	handler "github.com/noskine/pilot_api/pkg/handlers"
)

func main() {
	port := config.NewEnvironments().GetPort()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", handler.CreatePublicationsHandler)
	mux.HandleFunc("GET /api/publications", handler.GetAllPublicationHandler)
	mux.HandleFunc("GET /api/publications/by", handler.FindByIdPublicationHandler)

	log.Printf("Server is Runnign... Port => %s \n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
