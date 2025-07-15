package strategy

import (
	"log"
	"time"
)

type ParkingTicket struct {
	Plate        string
	CheckinDate  time.Time
	CheckoutDate time.Time
	Location     string
	Fare         float64
}

func NewParkingTicket(plate string, checkinDate time.Time, location string) *ParkingTicket {
	return &ParkingTicket{Plate: plate, CheckinDate: checkinDate, Location: location}
}

func (pt *ParkingTicket) Checkout(checkoutDate time.Time) {
	pt.CheckoutDate = checkoutDate
	fareCalculatorFactory := &FareCalculatorFactory{}
	fareCalculator, err := fareCalculatorFactory.Create(pt.Location)
	if err != nil {
		log.Fatal(err)
	}
	pt.Fare = fareCalculator.Calculate(pt.CheckinDate, checkoutDate)
}
