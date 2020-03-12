// Package car contains code for describing a car
package car

import (
	"encoding/json"
	"fmt"
)

// Car - interface wrapper for the car struct
type Car interface {
	String() string
	ToJSON() ([]byte, error)
}

type car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Color string `json:"color"`
	Year  int    `json:"year"`
}

func (c car) String() string {
	return fmt.Sprintln(c.Color, c.Year, c.Make, c.Model)
}

func (c car) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON - creates a car struct from JSON
func FromJSON(js string) (Car, error) {
	var c car
	err := json.Unmarshal([]byte(js), &c)
	return c, err
}
