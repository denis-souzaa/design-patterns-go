package chainofresponsability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustCalculateRideDuringNormalHours(t *testing.T) {

	overnightFare := &OverNightFareCalculator{}
	normalFare := &NormalFareCalculator{}
	normalFare.SetNext(overnightFare)
	sundayFare := &SundayFareCalculator{}
	overnightSundayFare := &OvernightSundayFareCalculator{}

	overnightFare.SetNext(sundayFare)
	sundayFare.SetNext(overnightSundayFare)
	r := NewRide(normalFare)

	r.AddSegment(10, time.Date(2021, 03, 01, 10, 0, 0, 0, time.Local))
	err := r.CalculateFare()
	assert.Nil(t, nil, err)
	assert.Equal(t, float64(21), r.Fare)
}

func TestMustCalculateRideOvernight(t *testing.T) {
	normalFare := &NormalFareCalculator{}
	overnightFare := &OverNightFareCalculator{}
	sundayFare := &SundayFareCalculator{}
	overnightSundayFare := &OvernightSundayFareCalculator{}

	sundayFare.SetNext(overnightSundayFare)
	overnightFare.SetNext(sundayFare)
	normalFare.SetNext(overnightFare)
	r := NewRide(normalFare)

	r.AddSegment(10, time.Date(2021, 03, 01, 23, 0, 0, 0, time.Local))
	err := r.CalculateFare()
	assert.Nil(t, nil, err)
	assert.Equal(t, float64(39), r.Fare)
}

func TestMustCalculateRideSunday(t *testing.T) {
	normalFare := &NormalFareCalculator{}
	overnightFare := &OverNightFareCalculator{}
	sundayFare := &SundayFareCalculator{}
	overnightSundayFare := &OvernightSundayFareCalculator{}

	sundayFare.SetNext(overnightSundayFare)
	overnightFare.SetNext(sundayFare)
	normalFare.SetNext(overnightFare)
	r := NewRide(normalFare)

	r.AddSegment(10, time.Date(2021, 03, 07, 10, 0, 0, 0, time.Local))
	err := r.CalculateFare()
	assert.Nil(t, nil, err)
	assert.Equal(t, float64(29), r.Fare)
}

func TestMustCalculateRideSundayNightly(t *testing.T) {
	normalFare := &NormalFareCalculator{}
	overnightFare := &OverNightFareCalculator{}
	sundayFare := &SundayFareCalculator{}
	overnightSundayFare := &OvernightSundayFareCalculator{}

	sundayFare.SetNext(overnightSundayFare)
	overnightFare.SetNext(sundayFare)
	normalFare.SetNext(overnightFare)
	r := NewRide(normalFare)

	r.AddSegment(10, time.Date(2021, 03, 07, 23, 0, 0, 0, time.Local))
	err := r.CalculateFare()
	assert.Nil(t, nil, err)
	assert.Equal(t, float64(50), r.Fare)
}

func TestNotMustCalculateRideWhenInvalidDistance(t *testing.T) {
	normalFare := &NormalFareCalculator{}
	overnightFare := &OverNightFareCalculator{}
	sundayFare := &SundayFareCalculator{}
	overnightSundayFare := &OvernightSundayFareCalculator{}

	normalFare.SetNext(overnightFare)
	overnightFare.SetNext(sundayFare)
	sundayFare.SetNext(overnightSundayFare)
	r := NewRide(normalFare)

	err := r.AddSegment(0, time.Date(2021, 03, 01, 23, 0, 0, 0, time.Local))
	assert.EqualError(t, err, "invalid distance")
}
