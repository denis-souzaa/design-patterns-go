package strategy

import (
	"denis-souzaa/design-patterns-go/config"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestMustCalculateFareAirport(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	plate := fmt.Sprintf("AAA%d", r.Intn(9000)+1000)
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	parkingTicketRepo := NewParkingTicketRepositoryDatabase(db)
	checkin := NewCheckin(parkingTicketRepo)
	inputCheckin := Input{
		plate:       plate,
		checkinDate: time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local),
		location:    "airport",
	}
	checkin.Execute(inputCheckin)
	checkout := NewCheckout(parkingTicketRepo)
	inputCheckout := InputCheckout{
		plate:        plate,
		checkoutDate: time.Date(2025, 07, 15, 12, 0, 0, 0, time.Local),
	}
	output := checkout.Execute(inputCheckout)
	assert.Equal(t, float64(20), output.fare)
}

func TestMustCalculateFareShopping(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	plate := fmt.Sprintf("AAA%d", r.Intn(9000)+1000)
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	parkingTicketRepo := NewParkingTicketRepositoryDatabase(db)
	checkin := NewCheckin(parkingTicketRepo)
	inputCheckin := Input{
		plate:       plate,
		checkinDate: time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local),
		location:    "shopping",
	}
	checkin.Execute(inputCheckin)
	checkout := NewCheckout(parkingTicketRepo)
	inputCheckout := InputCheckout{
		plate:        plate,
		checkoutDate: time.Date(2025, 07, 15, 15, 0, 0, 0, time.Local),
	}
	output := checkout.Execute(inputCheckout)
	assert.Equal(t, float64(30), output.fare)
}

func TestMustCalculateFareBeach(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	plate := fmt.Sprintf("AAA%d", r.Intn(9000)+1000)
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	parkingTicketRepo := NewParkingTicketRepositoryDatabase(db)
	checkin := NewCheckin(parkingTicketRepo)
	inputCheckin := Input{
		plate:       plate,
		checkinDate: time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local),
		location:    "beach",
	}
	checkin.Execute(inputCheckin)
	checkout := NewCheckout(parkingTicketRepo)
	inputCheckout := InputCheckout{
		plate:        plate,
		checkoutDate: time.Date(2025, 07, 15, 15, 0, 0, 0, time.Local),
	}
	output := checkout.Execute(inputCheckout)
	assert.Equal(t, float64(10), output.fare)
}
