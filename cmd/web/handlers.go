package main

import (
	"log"
	"net/http"
)

func (app *application) mainPage(w http.ResponseWriter, r *http.Request) {
	template, err := app.templateRepository.GetTemplate("index.gohtml")
	if err != nil {
		log.Panic(err)
	}

	template.Execute(w, nil)
}
