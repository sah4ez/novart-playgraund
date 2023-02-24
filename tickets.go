package main

import (
	"time"

	"github.com/sah4ez/novart-playgraund/pkg/types"
)

var t = types.Ticket{}

var (
	tickets = make(map[types.Ticket]float64, 0)
)

var (
	allTicketsID = []types.TicketID{
		"APPL",
		"AMZ",
	}
)

func InitData() {
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-24T00:00:00Z")
		ticket := types.Ticket{ID: "APPL", Date: ticketTime}
		tickets[ticket] = float64(1.0)
	}
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-23T00:00:00Z")
		ticket := types.Ticket{ID: "APPL", Date: ticketTime}
		tickets[ticket] = float64(2.0)
	}
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-22T00:00:00Z")
		ticket := types.Ticket{ID: "APPL", Date: ticketTime}
		tickets[ticket] = float64(3.0)
	}
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-24T00:00:00Z")
		ticket := types.Ticket{ID: "AMZ", Date: ticketTime}
		tickets[ticket] = float64(1.0)
	}
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-23T00:00:00Z")
		ticket := types.Ticket{ID: "AMZ", Date: ticketTime}
		tickets[ticket] = float64(2.0)
	}
	{
		ticketTime, _ := time.Parse(time.RFC3339, "2023-02-22T00:00:00Z")
		ticket := types.Ticket{ID: "AMZ", Date: ticketTime}
		tickets[ticket] = float64(3.0)
	}
}
