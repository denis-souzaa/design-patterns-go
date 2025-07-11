package builder

type FlightTickerBuilder struct {
	airline                   string
	flightCode                string
	fromAirport               string
	toAirport                 string
	passengerName             string
	passengerEmail            string
	passengerDocument         string
	passengerGender           string
	emergencyContactName      string
	emergencyContactTelephone string
	seat                      string
	checkedBags               int8
	hasCheckin                bool
	terminal                  string
	gate                      string
	priority                  int8
}

func (ft *FlightTickerBuilder) Flight(airline, flightCode string) *FlightTickerBuilder {
	ft.airline = airline
	ft.flightCode = flightCode
	return ft
}

func (ft *FlightTickerBuilder) Trip(from, to string) *FlightTickerBuilder {
	ft.fromAirport = from
	ft.toAirport = to
	return ft
}

func (ft *FlightTickerBuilder) Passsenger(name, email, document, gender string) *FlightTickerBuilder {
	ft.passengerName = name
	ft.passengerEmail = email
	ft.passengerDocument = document
	ft.passengerGender = gender
	return ft
}

func (ft *FlightTickerBuilder) EmergencyContact(name, telephone string) *FlightTickerBuilder {
	ft.emergencyContactName = name
	ft.emergencyContactTelephone = telephone
	return ft
}

func (ft *FlightTickerBuilder) Seat(seat string) *FlightTickerBuilder {
	ft.seat = seat
	return ft
}

func (ft *FlightTickerBuilder) CheckedBags(checkedBags int8) *FlightTickerBuilder {
	ft.checkedBags = checkedBags
	return ft
}

func (ft *FlightTickerBuilder) CheckeinInformation(hasCheckin bool, terminal, gate string) *FlightTickerBuilder {
	ft.hasCheckin = hasCheckin
	ft.terminal = terminal
	ft.gate = gate
	return ft
}

func (ft *FlightTickerBuilder) Priority(priority int8) *FlightTickerBuilder {
	ft.priority = priority
	return ft
}

func (ft *FlightTickerBuilder) FlightTicket() *FligthTicket {
	return NewFlightTicket(*ft)
}
