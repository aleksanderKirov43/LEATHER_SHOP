package models

import (
	"encoding/json"
)

type Category struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Props json.RawMessage `json:"props"`
}
