package models

import (
	"encoding/json"
)

type Category struct {
	Id    int             `json:"id" gorm:"primaryKey"`
	Name  string          `json:"name" gorm:"name"`
	Props json.RawMessage `json:"props"`
}
