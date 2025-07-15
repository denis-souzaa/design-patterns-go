package state

import (
	"errors"
)

type TicketStatus interface {
	Assign() error
	Start() error
	Close() error
	Status() string
}

type BaseStatus struct {
	value  string
	ticket *Ticket
}

type RequestedStatus struct {
	BaseStatus
}

func NewRequestedStatus(t *Ticket) *RequestedStatus {
	return &RequestedStatus{BaseStatus{value: "requested", ticket: t}}
}

func (rs *RequestedStatus) Assign() error {
	rs.ticket.status = NewAssignStatus(rs.ticket)
	return nil
}
func (rs *RequestedStatus) Start() error {
	return errors.New("could not start ticket")
}
func (rs *RequestedStatus) Close() error {
	return errors.New("could not close ticket")
}

func (rs *RequestedStatus) Status() string {
	return rs.value
}

type AssignStatus struct {
	BaseStatus
}

func NewAssignStatus(t *Ticket) *AssignStatus {
	return &AssignStatus{BaseStatus{value: "assigned", ticket: t}}
}

func (rs *AssignStatus) Assign() error {
	return errors.New("could not assign ticket")
}
func (rs *AssignStatus) Start() error {
	rs.ticket.status = NewInProgressStatus(rs.ticket)
	return nil
}
func (rs *AssignStatus) Close() error {
	return errors.New("could not close ticket")
}

func (rs *AssignStatus) Status() string {
	return rs.value
}

type InProgressStatus struct {
	BaseStatus
}

func NewInProgressStatus(t *Ticket) *InProgressStatus {
	return &InProgressStatus{BaseStatus{value: "in_progress", ticket: t}}
}

func (rs *InProgressStatus) Assign() error {
	return errors.New("could not assign ticket")
}
func (rs *InProgressStatus) Start() error {
	return errors.New("could not start ticket")
}
func (rs *InProgressStatus) Close() error {
	rs.ticket.status = NewClosedStatus(rs.ticket)
	return nil
}

func (rs *InProgressStatus) Status() string {
	return rs.value
}

type ClosedStatus struct {
	BaseStatus
}

func NewClosedStatus(t *Ticket) *ClosedStatus {
	return &ClosedStatus{BaseStatus{value: "closed", ticket: t}}
}

func (rs *ClosedStatus) Assign() error {
	return errors.New("could not assign ticket")
}
func (rs *ClosedStatus) Start() error {
	return errors.New("could not start ticket")
}
func (rs *ClosedStatus) Close() error {
	return errors.New("could not start ticket")
}

func (rs *ClosedStatus) Status() string {
	return rs.value
}
