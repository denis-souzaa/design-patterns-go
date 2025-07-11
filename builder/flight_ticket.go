package builder

type FligthTicket struct {
	Airline                   string
	FromAirport               string
	ToAirport                 string
	FlightCode                string
	PassengerName             string
	PassengerEmail            string
	PassengerDocument         string
	PassengerGender           string
	EmergencyContactName      string
	EmergencyContactTelephone string
	Seat                      string
	CheckedBags               int8
	HasCheckin                bool
	Terminal                  string
	Gate                      string
	Priority                  int8
}

func NewFlightTicket(b FlightTickerBuilder) *FligthTicket {
	return &FligthTicket{
		Airline:                   b.airline,
		FromAirport:               b.fromAirport,
		ToAirport:                 b.toAirport,
		FlightCode:                b.flightCode,
		PassengerName:             b.passengerName,
		PassengerEmail:            b.passengerEmail,
		PassengerDocument:         b.passengerDocument,
		PassengerGender:           b.passengerGender,
		EmergencyContactName:      b.emergencyContactName,
		EmergencyContactTelephone: b.emergencyContactTelephone,
		Seat:                      b.seat,
		CheckedBags:               b.checkedBags,
		HasCheckin:                b.hasCheckin,
		Terminal:                  b.terminal,
		Gate:                      b.gate,
		Priority:                  b.priority,
	}
}
