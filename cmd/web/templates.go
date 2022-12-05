package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear    int
	CurrentWeekDay string
}

func (app *application) addDefaultData(td *templateData, request *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.CurrentYear = time.Now().Year()
	td.CurrentWeekDay = time.Now().Weekday().String()

	return td
}

func newCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "html/pages/*.htm"))
	if err != nil {
		return nil, err
	}

	partials, err := filepath.Glob(filepath.Join(dir, "html/partials/*.htm"))
	if err != nil {
		return nil, err
	}

	base := []string{
		filepath.Join(dir, "html/layouts/base.htm"),
	}

	base = append(base, partials...)
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			page,
		}

		files = append(files, base...)
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
