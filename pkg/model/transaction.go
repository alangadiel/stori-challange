package model

import "time"

type Transaction struct {
	ID     int       `json:"id"`
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}
