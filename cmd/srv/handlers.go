package main

import (
	"attestation/pgk/models"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func (app *Application) ReadAllData(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("files/cities.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	var cities []models.City
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		cities = append(cities, models.City{
			ID:         line[0],
			Name:       line[1],
			Region:     line[2],
			District:   line[3],
			Population: line[4],
			Foundation: line[5],
		})
	}
	if err = json.NewEncoder(w).Encode(cities); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
