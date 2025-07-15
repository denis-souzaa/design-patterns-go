package strategy

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestMustCalculateTicketFare(t *testing.T) {
	parkingTicket := NewParkingTicket("AAA9999", time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local), "airport")
	parkingTicket.Checkout(time.Date(2025, 07, 15, 12, 0, 0, 0, time.Local))
	assert.Equal(t, float64(20), parkingTicket.Fare)
}
