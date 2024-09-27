package main

import (
	"log"
	"log/slog"
	"net/http"
)

type application struct {
	logger *slog.Logger
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting server on 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
