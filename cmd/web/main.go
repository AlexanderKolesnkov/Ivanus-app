package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type cfg struct {
	Addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	cfg := new(cfg)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "Сетевой адрес HTTP")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запусе веб-сервера на localhost%s", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
