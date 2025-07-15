package state

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestMustOpenTicket(t *testing.T) {
	customerId := 1
	ticket := NewTicket(customerId, time.Date(2025, 07, 15, 8, 0, 0, 0, time.Local))
	assert.Equal(t, "requested", ticket.Status())
	duration := ticket.Statistics(time.Date(2025, 07, 15, 9, 0, 0, 0, time.Local))
	assert.Equal(t, 1, duration.requestDuration)
	employeeId := 2
	ticket.Assign(employeeId, time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local))
	duration = ticket.Statistics(time.Date(2025, 07, 15, 11, 0, 0, 0, time.Local))
	assert.Equal(t, 1, duration.assignDuration)
	fmt.Println(ticket.Status())
	assert.Equal(t, "assigned", ticket.Status())
	ticket.Start(time.Date(2025, 07, 15, 16, 0, 0, 0, time.Local))
	assert.Equal(t, "in_progress", ticket.Status())
	ticket.Close(time.Date(2025, 07, 15, 18, 0, 0, 0, time.Local))
	assert.Equal(t, "closed", ticket.Status())
}
