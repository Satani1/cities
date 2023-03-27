package models

type City struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population string `json:"population"`
	Foundation string `json:"foundation"`
}
