package main

import (
	"net/http"
)

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		app.methodNotAllowedError(response)
		return
	}
	if request.URL.Path != "/" {
		app.notFoundError(response)
		return
	}

	snippets, err := app.dbConn.Latest()
	if err != nil {
		app.serverError(response, err)
	}

	td := &templateData{
		Snippets: snippets,
	}

	app.render(response, request, "home.htm", td)
}

func (app *application) write(response http.ResponseWriter, request *http.Request) {

}

func (app *application) post(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		app.methodNotAllowedError(response)
		return
	}
}

func (app *application) settings(response http.ResponseWriter, request *http.Request) {

}

func (app *application) logout(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		app.methodNotAllowedError(response)
		return
	}
}
