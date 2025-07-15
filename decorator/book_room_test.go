package main

import (
	"denis-souzaa/design-patterns-go/config"
	"denis-souzaa/design-patterns-go/decorator/booking"
	"denis-souzaa/design-patterns-go/decorator/room"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeCreateReservation(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	roomRepo := room.NewRepositoryDatabase(db)
	bookingRepo := booking.NewRepositoryDatabase(db)
	bookRoom := NewBookRoom(roomRepo, bookingRepo)
	input := Input{
		Email:        "john.doe@mail.com",
		CheckinDate:  time.Date(2026, 07, 11, 10, 0, 0, 0, time.Local),
		CheckoutDate: time.Date(2026, 07, 15, 10, 0, 0, 0, time.Local),
		Category:     "suite",
	}
	outBookRoom, _ := bookRoom.Execute(input)
	inputBooking := booking.Input{Code: outBookRoom.Code}
	bookingByCode := &booking.ByCode{Repo: bookingRepo}
	outBookingByCode := bookingByCode.Execute(inputBooking)
	assert.Equal(t, 4, outBookingByCode.Duration)
	assert.Equal(t, 2000.0, outBookingByCode.Price)
	cancelBooking := &booking.Cancel{Repo: bookingRepo}
	cancelBooking.Execute(inputBooking)
}
