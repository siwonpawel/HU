package main

import "net/http"

func (app *application) handlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.mainPage)

	return mux
}
