package strategy

import (
	"log"
	"time"
)

type InputCheckout struct {
	plate        string
	checkoutDate time.Time
}
type Output struct {
	plate string
	fare  float64
}

type Checkout struct {
	parkingTicketRepo ParkingTicketRepository
}

func NewCheckout(pTicketRepo ParkingTicketRepository) *Checkout {
	return &Checkout{parkingTicketRepo: pTicketRepo}
}

func (c *Checkout) Execute(i InputCheckout) Output {
	parkingTicket, err := c.parkingTicketRepo.ByPlate(i.plate)
	if err != nil {
		log.Fatal(err)
	}
	if parkingTicket.Plate == "" {
		log.Fatal("parking ticket not found")
	}
	parkingTicket.Checkout(i.checkoutDate)
	err = c.parkingTicketRepo.Update(*parkingTicket)
	if err != nil {
		log.Fatal(err)
	}
	return Output{plate: parkingTicket.Plate, fare: parkingTicket.Fare}
}
