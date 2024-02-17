package main

import (
	"log"
	"net/http"

	"github.com/siwonpawel/HU/ui"
)

func main() {
	files := ui.GetFS(false)

	templates, err := ui.NewTemplateRepository(files)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		template, err := templates.GetTemplate("index.gohtml")
		if err != nil {
			log.Panic(err)
		}

		template.Execute(w, nil)
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Panic(err)
	}
}
