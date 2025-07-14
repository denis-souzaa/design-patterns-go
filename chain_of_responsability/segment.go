package chainofresponsability

import (
	"errors"
	"time"
)

type Segment struct {
	distance float64
	date     time.Time
}

func NewSegment(distance float64, date time.Time) (*Segment, error) {

	if !isValidDistance(distance) {
		return nil, errors.New("invalid distance")
	}

	return &Segment{distance: distance, date: date}, nil
}

func isValidDistance(distance float64) bool {
	return distance != 0 && distance > 0
}

func (s *Segment) isOverNight() bool {
	return s.date.Hour() >= 22 || s.date.Hour() <= 6
}

func (s *Segment) isSunday() bool {
	return s.date.Weekday() == 0
}
