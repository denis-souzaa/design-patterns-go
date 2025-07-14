package chainofresponsability

import (
	"errors"
	"time"
)

type Ride struct {
	Segments       []Segment
	Fare           float64
	FareCalculator FareCalculator
}

func NewRide(fareCalculator FareCalculator) *Ride {
	return &Ride{FareCalculator: fareCalculator}
}

func (r *Ride) AddSegment(distance float64, date time.Time) error {
	s, err := NewSegment(distance, date)
	if err != nil {
		return errors.New("invalid distance")
	}
	r.Segments = append(r.Segments, *s)
	return nil
}

func (r *Ride) CalculateFare() error {
	var fare float64
	for _, s := range r.Segments {
		fare += r.FareCalculator.Calculate(s)
	}

	if fare < 10 {
		r.Fare = 10
	}
	r.Fare = fare
	return nil
}
