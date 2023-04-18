package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddress = ":3030"

func main() {
	fmt.Println("Server running in ", httpAddress)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(httpAddress, mux)

	log.Println(err)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Todo bien"))
	if err != nil {
		return
	}
}
