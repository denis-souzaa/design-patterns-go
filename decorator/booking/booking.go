package booking

import (
	"denis-souzaa/design-patterns-go/decorator/room"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Code         string
	RoomId       int
	Email        string
	CheckinDate  time.Time
	CheckoutDate time.Time
	Duration     int
	Price        float64
	status       string
}

func (b *Booking) Create(email string, room room.Room, checkinDate, checkoutDate time.Time) *Booking {
	code := uuid.NewString()
	duration := int(checkoutDate.Day() - checkinDate.Day())
	price := float64(duration * int(room.Price))
	status := "confirmed"

	return &Booking{
		Code:         code,
		RoomId:       room.RoomId,
		Email:        email,
		CheckinDate:  checkinDate,
		CheckoutDate: checkoutDate,
		Duration:     duration,
		Price:        price,
		status:       status,
	}
}

func (b *Booking) Cancel() {
	b.status = "canceled"
}
