package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	cache   map[string]*template.Template
}

func main() {
	addr := flag.Int("addr", 4000, "network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cache, err := newCache("./ui/")
	if err != nil {
		errLog.Fatal(err)
	}

	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
		cache:   cache,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", *addr),
		Handler:      app.routes(),
		ErrorLog:     errLog,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting the Server on port %d\n", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errLog.Fatal(err)
	}

}
