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

	rMux.HandleFunc("/city/{id:[0-9999]+}", app.GetInfo).Methods("GET")
	rMux.HandleFunc("/create", app.CreateCity).Methods("POST")
	rMux.HandleFunc("/delete", app.DeleteCity).Methods("DELETE")
	rMux.HandleFunc("/update-population", app.UpdatePopulation).Methods("PUT")
	rMux.HandleFunc("/list-reg", app.ListByRegion).Methods("GET")
	rMux.HandleFunc("/list-dist", app.ListByDistrict).Methods("GET")
	rMux.HandleFunc("/list-pop", app.ListByPopulation).Methods("GET")
	rMux.HandleFunc("/list-found", app.ListByFoundation).Methods("GET")
	rMux.HandleFunc("/file", app.testWrite).Methods("GET")

	return rMux
}
