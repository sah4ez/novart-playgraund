package main

import (
	"fmt"
	"math"
	"time"

	"github.com/sah4ez/novart-playgraund/pkg/types"
)

type (
	Pricer interface {
		LoadPrice(id types.TicketID, date time.Time) (price float64, ok bool)
	}
)

func main() {
	InitData()

	var pricer Pricer
	pricer = &ServicePricer{}

	ticker := time.NewTicker(1 * time.Second)
	startDate, _ := time.Parse(time.RFC3339, "2023-02-22T00:00:00Z")
	done := make(chan struct{})
	for {
		nextDate := startDate.Add(24 * time.Hour)
		select {
		case <-ticker.C:
			for _, ticket := range allTicketsID {
				t1, _ := pricer.LoadPrice(ticket, startDate)
				t2, ok := pricer.LoadPrice(ticket, nextDate)
				if !ok {
					fmt.Println("all currencies check")
					done <- struct{}{}
					ticker.Stop()
					break
				}

				{
					ok := compare(t1, t2)
					if ok {
						fmt.Println(startDate.Format(time.RFC3339), ticket, "ok", t1, t2)
					} else {
						fmt.Println(startDate.Format(time.RFC3339), ticket, "not ok", t1, t2)
					}
				}
			}
		case <-done:
			break
		}
		startDate = nextDate
	}
}

func compare(ticket1 float64, ticket2 float64) (ok bool) {
	percent := 100.0 - (ticket2 * 100.0 / ticket1)
	absPercetn := math.Abs(percent)
	return absPercetn >= 1.0
}
