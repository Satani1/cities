package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func (app *Application) readData() {

	file, err := os.Open("files/cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
