package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// SETUP PORT
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// LOGGING
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// HTTP SERVER
	mux := http.NewServeMux()

	// FILE SERVER
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// ROUTING
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/home", app.snippetView)
	mux.HandleFunc("/snipper/create", app.snippetCreate)

	// SERVER CUSTOM
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
