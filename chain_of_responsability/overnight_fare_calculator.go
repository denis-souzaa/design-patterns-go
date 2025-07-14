package chainofresponsability

const (
	OVERNIGHT_FARE = 3.9
)

type OverNightFareCalculator struct {
	Next FareCalculator
}

func (oc *OverNightFareCalculator) Calculate(s Segment) float64 {
	if s.isOverNight() && !s.isSunday() {
		return s.distance * OVERNIGHT_FARE
	}
	return oc.Next.Calculate(s)
}

func (oc *OverNightFareCalculator) SetNext(next FareCalculator) {
	oc.Next = next
}
