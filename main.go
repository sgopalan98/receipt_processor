package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sgoplan98/receipt_processor/api/handlers"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}

func main() {

	// TODO: Does this execute concurrently? Need to research and change code according to that!
	// TODO: What is mux?
	// mux := http.NewServeMux()
	// mux.HandleFunc("/receipts/process", http.HandlerFunc(receiptsProcessHandler))
	// mux.HandleFunc("/receipts", http.HandlerFunc(receiptsPointsHandler))
	// mux.HandleFunc("/", http.HandlerFunc(indexHandler))

	http.HandleFunc("/receipts/process", handlers.ReceiptsProcessHandler)
	http.HandleFunc("/receipts/", handlers.ReceiptsPointsHandler)
	http.HandleFunc("/", indexHandler)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
