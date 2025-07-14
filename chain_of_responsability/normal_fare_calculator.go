package chainofresponsability

const (
	FARE = 2.10
)

type NormalFareCalculator struct {
	Next FareCalculator
}

func (nc *NormalFareCalculator) Calculate(s Segment) float64 {

	if !s.isOverNight() && !s.isSunday() {
		return s.distance * FARE
	}
	return nc.Next.Calculate(s)
}

func (nc *NormalFareCalculator) SetNext(next FareCalculator) {
	nc.Next = next
}
