package strategy

import (
	"database/sql"
	"denis-souzaa/design-patterns-go/config"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestMustCalculateFare(t *testing.T) {

	generatedCarPlate := func(t *testing.T) string {
		t.Helper()
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		return fmt.Sprintf("AAA%d", r.Intn(9000)+1000)
	}

	connectionDb := func(t *testing.T) *sql.DB {
		db, err := config.New()
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

	t.Run("Deve calcular tarifa para aeroporto", func(t *testing.T) {
		plate := generatedCarPlate(t)
		db := connectionDb(t)
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
	})

	t.Run("Deve calcular tarifa para shopping", func(t *testing.T) {
		plate := generatedCarPlate(t)
		db := connectionDb(t)
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
	})

	t.Run("Deve calcular tarifa para a praia", func(t *testing.T) {
		plate := generatedCarPlate(t)
		db := connectionDb(t)
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
	})
}
