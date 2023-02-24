package main

import (
	"time"

	"github.com/sah4ez/novart-playgraund/pkg/types"
)

type ServicePricer struct{}

func (s *ServicePricer) LoadPrice(id types.TicketID, date time.Time) (price float64, ok bool) {

	ticket := types.Ticket{
		ID:   id,
		Date: date,
	}
	price, ok = tickets[ticket]
	return
}
