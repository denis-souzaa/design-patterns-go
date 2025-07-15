package state

import (
	"time"
)

type Duration struct {
	requestDuration int
	assignDuration  int
}

type Ticket struct {
	customerId  int
	status      TicketStatus
	requestDate time.Time
	employeeId  int
	assignDate  time.Time
	startDate   time.Time
	endDate     time.Time
}

func NewTicket(customerId int, requestDate time.Time) *Ticket {
	ticket := &Ticket{customerId: customerId, requestDate: requestDate}
	ticket.status = NewRequestedStatus(ticket)
	return ticket
}

func (t *Ticket) Assign(employeeId int, assignedDate time.Time) {
	t.status.Assign()
	t.employeeId = employeeId
	t.assignDate = assignedDate
}

func (t *Ticket) Start(startDate time.Time) {
	t.status.Start()
	t.startDate = startDate
}

func (t *Ticket) Close(endDate time.Time) {
	t.status.Close()
	t.endDate = endDate
}

func (t *Ticket) Statistics(currentDate time.Time) Duration {
	var requestDuration, assignDuration int

	if t.assignDate.IsZero() {
		requestDuration = currentDate.Hour() - t.requestDate.Hour()
	} else {
		requestDuration = t.requestDate.Hour() - t.assignDate.Hour()
	}

	if t.startDate.IsZero() {
		assignDuration = currentDate.Hour() - t.assignDate.Hour()
	} else {
		assignDuration = t.startDate.Hour() - t.assignDate.Hour()
	}

	return Duration{requestDuration: requestDuration, assignDuration: assignDuration}
}

func (t *Ticket) Status() string {
	return t.status.Status()
}
