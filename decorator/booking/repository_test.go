package booking

import (
	"denis-souzaa/design-patterns-go/decorator/config"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldSaveBooking(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	b := Booking{
		Code:         "abc",
		RoomId:       1,
		Email:        "johh.doe@mail.com",
		CheckinDate:  time.Now().AddDate(2025, 07, 10),
		CheckoutDate: time.Now().AddDate(2025, 07, 12),
		Duration:     2,
		Price:        1000,
		status:       "confirmed",
	}
	repo := NewRepositoryDatabase(db)
	err = repo.Save(b)
	assert.Equal(t, nil, err)
}
