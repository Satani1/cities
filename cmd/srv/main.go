package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	Addr     string
}

func main() {
	//addr config from terminal
	addr := flag.String("addr", "localhost:4000", "Server Address")
	flag.Parse()

	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERORR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//srv model
	App := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		Addr:     *addr,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  App.Routes(),
	}
	//launch
	infoLog.Printf("Launching server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
