package main

import (
	"net/http"

	"github.com/grepex/dead-mans-switch/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// TODO: Add distinction between protected/unprotected routes

	mux.HandleFunc("GET /{$}", app.home)

	return mux
}
