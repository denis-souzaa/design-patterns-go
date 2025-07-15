package strategy

import (
	"log"
	"time"
)

type Input struct {
	plate       string
	checkinDate time.Time
	location    string
}
type Checkin struct {
	parkingTicketRepo ParkingTicketRepository
}

func NewCheckin(pTicketRepo ParkingTicketRepository) *Checkin {
	return &Checkin{parkingTicketRepo: pTicketRepo}
}

func (c *Checkin) Execute(i Input) {
	existingParkingTicket, err := c.parkingTicketRepo.ByPlate(i.plate)
	if err != nil {
		log.Fatal(err)
	}
	if existingParkingTicket.Plate != "" {
		log.Fatal("duplicated plate")
	}

	parkingTicket := NewParkingTicket(i.plate, i.checkinDate, i.location)
	c.parkingTicketRepo.Save(*parkingTicket)
}
