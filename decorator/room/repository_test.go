package room

import (
	"denis-souzaa/design-patterns-go/config"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetRoom(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewRepositoryDatabase(db)
	room, _ := repo.ById(1)
	assert.Equal(t, "suite", room.Category)
	assert.Equal(t, 500.0, room.Price)
}

func TestShouldGetRoomAvailable(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewRepositoryDatabase(db)
	rooms, _ := repo.AvailableRoomsByPeriodAndCategory(time.Date(2025, 07, 12, 10, 0, 0, 0, time.Local), time.Date(2025, 07, 12, 10, 0, 0, 0, time.Local), "suite")
	for _, r := range rooms {
		assert.Equal(t, "suite", r.Category)
		assert.Equal(t, 500.0, r.Price)
	}
}

func TestShouldGetRoomAvailableNotReturnRoomAvailable(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewRepositoryDatabase(db)
	rooms, _ := repo.AvailableRoomsByPeriodAndCategory(time.Date(2025, 07, 11, 10, 0, 0, 0, time.Local), time.Date(2025, 07, 15, 10, 0, 0, 0, time.Local), "suite")
	assert.Len(t, rooms, 0)
}
