package ui

import (
	"fmt"
	"html/template"
	"io/fs"
)

type TemplateRepository interface {
	GetTemplate(name string) (*template.Template, error)
}

func NewCachedTemplateRepository() (TemplateRepository, error) {
	templates, err := loadTemplates(getFS(true))
	if err != nil {
		return nil, err
	}

	return &cachedTemplateRepository{templates: templates}, nil
}

type cachedTemplateRepository struct {
	templates map[string]*template.Template
}

func (tr *cachedTemplateRepository) GetTemplate(name string) (*template.Template, error) {
	tmpl, ok := tr.templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}

	return tmpl, nil
}

func NewTemplateRepository() (TemplateRepository, error) {
	files := getFS(false)
	_, err := loadTemplates(files)
	if err != nil {
		return nil, err
	}

	return &templateRepository{files: files}, nil
}

type templateRepository struct {
	files fs.FS
}

func (tr *templateRepository) GetTemplate(name string) (*template.Template, error) {
	templates, err := loadTemplates(tr.files)
	if err != nil {
		return nil, err
	}

	tmpl, ok := templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}

	return tmpl, nil
}
