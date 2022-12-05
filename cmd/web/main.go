package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"beyondthemataulghuroor.com/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	cache   map[string]*template.Template
	dbConn  *models.Conn
}

func main() {
	addr := flag.Int("addr", 4000, "network address")
	dsn := flag.String("dsn", os.Getenv("BTMG_DB_DSN"), "data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cache, err := newCache("./ui/")
	if err != nil {
		errLog.Fatal(err)
	}

	db, err := openDB("mysql", *dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Println("Database connection established successfully")

	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
		cache:   cache,
		dbConn: &models.Conn{
			DB: db,
		},
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

func openDB(name string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(name, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
