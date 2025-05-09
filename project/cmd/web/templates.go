package main

import (
	"html/template"
  "path/filepath"
	"time"
	"io/fs"
	"github.com/Jojaror/mi_primer_api_en_go/project/internal/models"
	"github.com/Jojaror/mi_primer_api_en_go/project/ui"
)

type templateData struct {
	CurrentYear			int
	Snippet 				*models.Snippet
	Snippets 				[]*models.Snippet
	Form 						any
	Flash 					string
	IsAuthenticated	bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap {
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
