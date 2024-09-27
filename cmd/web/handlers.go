package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homepage"))
}

func (app *application) messageView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This will render a message that has been scheduled"))
}

func (app *application) messageCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This will display a form to create and schedule a message"))
}

func (app *application) messageCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This will send the POST request"))
}
