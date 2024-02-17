package main

import (
	"net/http"

	"github.com/siwonpawel/HU/ui"
)

type application struct {
	addr               string
	templateRepository ui.TemplateRepository
}

func (a *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handlers().ServeHTTP(w, r)
}
