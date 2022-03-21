package main

import (
	"02-JSON-demo/handler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	handler.Start = time.Now()

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(handler.COVIDCASES, handler.DiagHandler)
	http.HandleFunc(handler.COVIDPOLICY, handler.UniInfoHandler)
	http.HandleFunc(handler.COVIDSTATUS, handler.NeighbourUniHandler)
	http.HandleFunc(handler.COVIDNOTIFICATIONS, handler.NeighbourUniHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
