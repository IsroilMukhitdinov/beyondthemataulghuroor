package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", app.home)
	router.HandleFunc("/write", app.write)
	router.HandleFunc("/post", app.post)
	router.HandleFunc("/settings", app.settings)
	router.HandleFunc("/logout", app.logout)

	return app.logRequest(router)
}
