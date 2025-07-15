package strategy

import (
	"errors"
	"time"
)

type FareCalculator interface {
	Calculate(checkinDate, checkoutDate time.Time) float64
}

type AirportFareCalculator struct{}

func (afc *AirportFareCalculator) Calculate(checkinDate, checkoutDate time.Time) float64 {
	diff := checkoutDate.Hour() - checkinDate.Hour()
	return float64(diff * 10)
}

type ShoppingFareCalculator struct{}

func (sfc *ShoppingFareCalculator) Calculate(checkinDate, checkoutDate time.Time) float64 {
	fare := float64(10)
	diff := checkoutDate.Hour() - checkinDate.Hour()
	remainingHours := diff - 3
	if remainingHours > 0 {
		fare += float64(remainingHours * 10)
	}
	return fare
}

type BeachFareCalculator struct{}

func (bfc *BeachFareCalculator) Calculate(checkinDate, checkoutDate time.Time) float64 {
	return 10
}

type FareCalculatorFactory struct{}

func (fcf *FareCalculatorFactory) Create(location string) (FareCalculator, error) {
	if location == "airport" {
		return &AirportFareCalculator{}, nil
	}
	if location == "shopping" {
		return &ShoppingFareCalculator{}, nil
	}
	if location == "beach" {
		return &BeachFareCalculator{}, nil
	}
	return nil, errors.New("location not found")
}
