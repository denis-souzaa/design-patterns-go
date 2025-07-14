package chainofresponsability

type FareCalculator interface {
	SetNext(FareCalculator)
	Calculate(segment Segment) float64
}
