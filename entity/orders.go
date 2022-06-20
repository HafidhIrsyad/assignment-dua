package entity

import "time"

type Orders struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Items   `json:"items"`
}
