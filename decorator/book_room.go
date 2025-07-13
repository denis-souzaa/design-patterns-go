package main

import (
	"denis-souzaa/design-patterns-go/decorator/booking"
	"denis-souzaa/design-patterns-go/decorator/room"
	"errors"
	"log"
	"time"
)

type BookRoom struct {
	RoomRepo    room.Repository
	BookingRepo booking.Repository
}

func NewBookRoom(roomRepo room.Repository, bookingRepo booking.Repository) *BookRoom {
	return &BookRoom{RoomRepo: roomRepo, BookingRepo: bookingRepo}
}

func (br *BookRoom) Execute(i Input) (*Output, error) {
	rooms, err := br.RoomRepo.AvailableRoomsByPeriodAndCategory(i.CheckinDate, i.CheckoutDate, i.Category)
	if err != nil {
		log.Fatal(err)
	}

	if len(rooms) == 0 {
		return nil, errors.New("room is not available")
	}

	var booking *booking.Booking
	booking = booking.Create(i.Email, *rooms[0], i.CheckinDate, i.CheckoutDate)
	err = br.BookingRepo.Save(*booking)
	if err != nil {
		log.Fatal(err)
	}
	return &Output{Code: booking.Code}, nil
}

type Input struct {
	Email        string
	CheckinDate  time.Time
	CheckoutDate time.Time
	Category     string
}

type Output struct {
	Code string
}
