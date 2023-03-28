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
