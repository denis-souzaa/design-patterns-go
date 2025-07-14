package chainofresponsability

const (
	OVERNIGHT_SUNDAY_FARE = 5
)

type OvernightSundayFareCalculator struct {
	Next FareCalculator
}

func (nc *OvernightSundayFareCalculator) Calculate(s Segment) float64 {
	if s.isOverNight() && s.isSunday() {
		return s.distance * OVERNIGHT_SUNDAY_FARE
	}

	return nc.Next.Calculate(s)
}

func (nc *OvernightSundayFareCalculator) SetNext(next FareCalculator) {
	nc.Next = next
}
