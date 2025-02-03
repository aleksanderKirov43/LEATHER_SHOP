package models

type Products struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Image       string `json:"image"`
	Sale        int    `json:"sale"`
	Price       int    `json:"price"`
	Category    int    `json:"category"`
	Property    int    `json:"property"`
}
