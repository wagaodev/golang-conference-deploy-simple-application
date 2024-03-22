package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", basePath)
	http.ListenAndServe(":8080", mux)
}

func basePath(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - Deploying application"))
}
