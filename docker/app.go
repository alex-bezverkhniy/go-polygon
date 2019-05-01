package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello from go!`))
}

// Simple HTTP server.
func main() {
	Mux := http.NewServeMux()
	Mux.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8081", Mux))
}
