package main

import (
	"denis-souzaa/design-patterns-go/decorator/booking"
	"denis-souzaa/design-patterns-go/decorator/config"
	"denis-souzaa/design-patterns-go/decorator/room"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldImportListBookings(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	roomRepo := room.NewRepositoryDatabase(db)
	bookingRepo := booking.NewRepositoryDatabase(db)

	input := ImportInput{
		File: `email;checkin_date;checkout_date;category;
		john.doe1@mail.com;2025-07-10 10:00:00;2025-07-12 10:00:00;suite;
		john.doe2@mail.com;2025-07-13 10:00:00;2025-07-15 10:00:00;suite;
		john.doe3@mail.com;2025-07-20 10:00:00;2025-07-22 10:00:00;suite;`,
	}
	importBooking := NewImportBooking(NewBookRoom(roomRepo, bookingRepo))
	outImportBooking, _ := importBooking.Execute(input)

	for _, c := range outImportBooking {
		inputBooking := booking.Input{Code: c}
		bookingByCode := booking.ByCode{Repo: bookingRepo}
		outBookingByCode := bookingByCode.Execute(inputBooking)
		assert.Equal(t, 2, outBookingByCode.Duration)
		assert.Equal(t, 1000.0, outBookingByCode.Price)
		cancelBooking := booking.Cancel{Repo: bookingRepo}
		cancelBooking.Execute(inputBooking)
	}
}
