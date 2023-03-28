package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Application) Routes() *mux.Router {
	rMux := mux.NewRouter()

	rMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addrResp := app.Addr
		if _, err := w.Write([]byte("Hello, im service for attestation. My addres is - " + addrResp)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	rMux.HandleFunc("/data", app.ReadAllData)
	rMux.HandleFunc("/city/{id:[0-9999]+}", app.GetInfo)
	rMux.HandleFunc("/test", app.testCity)

	return rMux
}
