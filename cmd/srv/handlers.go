package main

import (
	"attestation/pgk/models"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (app *Application) ReadAllData(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("files/cities.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		district := json.Number(line[4])
		foundation := json.Number(line[5])
		app.DataDB[id] = models.City{
			Name:       line[1],
			Region:     line[2],
			District:   line[3],
			Population: district,
			Foundation: foundation,
		}
	}
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(app.DataDB); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *Application) testCity(w http.ResponseWriter, r *http.Request) {
	var c models.City = app.DataDB[490]
	fmt.Fprintf(w, "City: %v\n", c)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (app *Application) GetInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var c models.City = app.DataDB[ID]
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
