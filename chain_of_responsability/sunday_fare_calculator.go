package chainofresponsability

const (
	SUNDAY_FARE = 2.9
)

type SundayFareCalculator struct {
	Next FareCalculator
}

func (nc *SundayFareCalculator) Calculate(s Segment) float64 {
	if !s.isOverNight() && s.isSunday() {
		return s.distance * SUNDAY_FARE
	}

	return nc.Next.Calculate(s)
}

func (nc *SundayFareCalculator) SetNext(next FareCalculator) {
	nc.Next = next
}
