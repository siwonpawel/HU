package ui

import (
	"html/template"
	"io/fs"
	"path"
)

func loadTemplates(files fs.FS) (map[string]*template.Template, error) {
	_, err := files.Open("html/base.gohtml")
	if err != nil {
		return nil, err
	}

	pages, err := fs.Glob(files, "html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)
	for _, page := range pages {
		name := path.Base(page)

		patterns := []string{
			"html/base.gohtml",
			"html/partials/*.gohtml",
			page,
		}

		template, err := template.New("base").ParseFS(files, patterns...)
		if err != nil {
			return nil, err
		}

		templates[name] = template
	}

	return templates, nil
}
