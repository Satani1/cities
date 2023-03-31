package main

import (
	"attestation/pgk/models"
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
	DataDB   map[int]models.City
}

func NewApp() *Application {
	//addr config from terminal
	addr := flag.String("addr", "localhost:4000", "Server Address")
	flag.Parse()
	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		Addr:     *addr,
		DataDB:   map[int]models.City{},
	}
	return App
}
func main() {

	App := NewApp()

	srv := &http.Server{
		Addr:     App.Addr,
		ErrorLog: App.errorLog,
		Handler:  App.Routes(),
	}

	App.ReadAndCashFileData()
	//launch
	App.infoLog.Printf("Launching server on %s", App.Addr)
	err := srv.ListenAndServe()
	App.errorLog.Fatal(err)
}
