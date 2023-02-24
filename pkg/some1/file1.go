package some1

import "fmt"

type PackageTicket struct {
	TicketID string
}

var GlogbalTicket = PackageTicket{TicketID: "123"}

func Print() {
	fmt.Println("hello world")
}
