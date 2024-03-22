package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", basePath)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error when starting server: %s", err)
	}
}

func basePath(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - Deploying application"))
}
