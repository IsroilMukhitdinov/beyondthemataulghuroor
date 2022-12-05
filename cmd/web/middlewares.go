package main

import "net/http"

func (app *application) logRequest(next http.Handler) http.Handler {
	fn := func(response http.ResponseWriter, request *http.Request) {
		app.infoLog.Printf("%s - %s", request.RemoteAddr, request.URL.Path)

		next.ServeHTTP(response, request)
	}

	return http.HandlerFunc(fn)
}
