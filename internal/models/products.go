package models

import "github.com/lib/pq"

type Products struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	Image       pq.StringArray `json:"image"`
	Sale        int            `json:"sale"`
	Price       int            `json:"price"`
	Status      int            `json:"status"`
	Category    int            `json:"category"`
	Property    int            `json:"property"`
}
