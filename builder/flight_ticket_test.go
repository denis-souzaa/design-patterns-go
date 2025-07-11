package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewFlightTicket(t *testing.T) {
	b := FlightTickerBuilder{}
	b.Flight("Azul", "9876").
		Trip("PMW", "GRU").
		Passsenger("John Doe", "john.doe@mail.com", "111.111.111-11", "M").
		EmergencyContact("Bob Simpson", "551111111-1111").
		Seat("4A").
		CheckeinInformation(true, "4A", "4").
		Priority(5)
	ticket := NewFlightTicket(b)
	assert.Equal(t, "John Doe", ticket.PassengerName)
}
