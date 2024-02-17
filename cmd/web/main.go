package main

import (
	"log"
	"net/http"

	"github.com/siwonpawel/HU/ui"
)

func main() {
	templateRepository, err := ui.NewTemplateRepository()
	if err != nil {
		log.Panic(err)
	}

	app := &application{
		templateRepository: templateRepository,
	}

	http.ListenAndServe("localhost:8080", app)
}
