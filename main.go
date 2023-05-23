package main

import (
	"log"
	"net/http"

	"github.com/sgoplan98/receipt_processor/api/handlers"
)

func main() {

	// TODO: Does this execute concurrently? Need to research and change code according to that!
	http.HandleFunc("/receipts/process", handlers.ReceiptsProcessHandler)
	http.HandleFunc("/receipts/", handlers.ReceiptsPointsHandler)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
