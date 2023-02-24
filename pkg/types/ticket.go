package types

import "time"

type TicketID string

type Ticket struct {
	ID    TicketID  `json:"id"`
	Date  time.Time `json:"date"`
	Price float64   `json:"price"`
}
