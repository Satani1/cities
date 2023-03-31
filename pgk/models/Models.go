package models

import "encoding/json"

type City struct {
	//ID         json.Number `json:"id"`
	Name       string      `json:"name"`
	Region     string      `json:"region"`
	District   string      `json:"district"`
	Population json.Number `json:"population"`
	Foundation json.Number `json:"foundation"`
}

type CityID struct {
	ID json.Number `json:"id"`
}

type PopulationID struct {
	ID         json.Number `json:"id"`
	Population json.Number `json:"population"`
}

type CityRange struct {
	PopulationStart json.Number `json:"populationStart"`
	PopulationEnd   json.Number `json:"populationEnd"`
	FoundationStart json.Number `json:"foundationStart"`
	FoundationEnd   json.Number `json:"foundationEnd"`
}
