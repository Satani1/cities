package main

import (
	"attestation/pgk/models"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func (app *Application) ReadAndCashFileData() {
	file, err := os.Open("files/cities.csv")
	if err != nil {
		app.errorLog.Fatal(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			app.errorLog.Fatal(err)
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			app.errorLog.Fatal(err)
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
}
func (app *Application) RewriteFileData() {
	//open file
	csvFile, err := os.Create("files/cities.csv")
	if err != nil {
		app.errorLog.Fatal(err)
	}
	defer csvFile.Close()
	//writer
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Flush()
	//slice of entries for writing
	records := [][]string{}

	for key, value := range app.DataDB {
		//convert json.numbers to int64 then convert it to string
		nPopulation, err := value.Population.Int64()
		Population := strconv.Itoa(int(nPopulation))
		if err != nil {
			app.errorLog.Fatal(err)
		}
		nFoundation, err := value.Population.Int64()
		Foundation := strconv.Itoa(int(nFoundation))
		if err != nil {
			app.errorLog.Fatal(err)
		}
		//create slice with data of city
		tempRec := []string{strconv.Itoa(key), value.Name, value.Region, value.District, Population, Foundation}
		//append data slice to slice of entries
		records = append(records, tempRec)
		fmt.Println(records)
	}
	//write records in file
	for _, record := range records {
		fmt.Println(record)
		err := csvWriter.Write(record)
		if err != nil {
			app.errorLog.Fatalln(err)
		}
		app.infoLog.Fatalln("Records to write in this session : ", record)
	}
}

func (app *Application) testWrite(w http.ResponseWriter, r *http.Request) {
	app.RewriteFileData()
	w.Write([]byte("test writing data to file :/"))
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
	var c models.City
	if _, ok := app.DataDB[ID]; ok {
		c = app.DataDB[ID]
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (app *Application) CreateCity(w http.ResponseWriter, r *http.Request) {
	var c models.City

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	min, max := 2000, 9999
	id := rand.Intn(max-min) + min
	app.DataDB[id] = c

	if err := json.NewEncoder(w).Encode(app.DataDB[id]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("Created city with an id - " + strconv.Itoa(id))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)

}

func (app *Application) DeleteCity(w http.ResponseWriter, r *http.Request) {
	var i models.CityID

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ID, err := i.ID.Int64()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := app.DataDB[int(ID)]; ok {
		delete(app.DataDB, int(ID))
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Application) UpdatePopulation(w http.ResponseWriter, r *http.Request) {
	var np models.PopulationID

	if err := json.NewDecoder(r.Body).Decode(&np); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := np.ID.Int64()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	nID := int(id)
	if entry, ok := app.DataDB[nID]; ok {
		entry.Population = np.Population

		app.DataDB[nID] = entry
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	if err := json.NewEncoder(w).Encode(app.DataDB[nID]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (app *Application) ListByRegion(w http.ResponseWriter, r *http.Request) {
	var reg models.City

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var records []models.City
	for _, entry := range app.DataDB {
		if entry.Region == reg.Region {
			records = append(records, entry)
		}
	}
	if len(records) == 0 {
		http.Error(w, "Not Found any cities in this region ;(", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (app *Application) ListByDistrict(w http.ResponseWriter, r *http.Request) {
	var reg models.City

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var records []models.City
	for _, entry := range app.DataDB {
		if entry.District == reg.District {
			records = append(records, entry)
		}
	}
	if len(records) == 0 {
		http.Error(w, "Not Found any cities in this district :(", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *Application) ListByPopulation(w http.ResponseWriter, r *http.Request) {
	var rangePop models.CityRange

	if err := json.NewDecoder(r.Body).Decode(&rangePop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var records []models.City
	for _, entry := range app.DataDB {
		if entry.Population >= rangePop.PopulationStart && entry.Population <= rangePop.PopulationEnd {
			records = append(records, entry)
		}
	}

	if len(records) == 0 {
		http.Error(w, "Not Found any city in that population range 0_o", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Application) ListByFoundation(w http.ResponseWriter, r *http.Request) {
	var rangePop models.CityRange

	if err := json.NewDecoder(r.Body).Decode(&rangePop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var records []models.City
	for _, entry := range app.DataDB {
		if entry.Foundation >= rangePop.FoundationStart && entry.Foundation <= rangePop.FoundationEnd {
			records = append(records, entry)
		}
	}

	if len(records) == 0 {
		http.Error(w, "Not Found any city in that foundation date range 0_o", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
