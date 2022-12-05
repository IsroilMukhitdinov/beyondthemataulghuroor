package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(response http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
	http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(response http.ResponseWriter, status int) {
	http.Error(response, http.StatusText(status), status)
}

func (app *application) notFoundError(response http.ResponseWriter) {
	app.clientError(response, http.StatusNotFound)
}

func (app *application) methodNotAllowedError(response http.ResponseWriter) {
	app.clientError(response, http.StatusMethodNotAllowed)
}
