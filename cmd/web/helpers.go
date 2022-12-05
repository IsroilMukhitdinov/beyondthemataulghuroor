package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(response http.ResponseWriter, request *http.Request, name string, td *templateData) {
	ts, ok := app.cache[name]
	if !ok {
		app.serverError(response, fmt.Errorf("template %s does not exist", name))
		return
	}
	td = app.addDefaultData(td, request)
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(response, err)
		return
	}

	buf.WriteTo(response)
}
