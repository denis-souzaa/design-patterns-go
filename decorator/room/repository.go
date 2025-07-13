package room

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Repository interface {
	AvailableRoomsByPeriodAndCategory(checkinDate, checkoutDate time.Time, category string) ([]*Room, error)
	ById(roomId int) (*Room, error)
}

type RepositoryDatabase struct {
	db *sql.DB
}

func NewRepositoryDatabase(db *sql.DB) *RepositoryDatabase {
	return &RepositoryDatabase{
		db: db,
	}
}

func (rd *RepositoryDatabase) AvailableRoomsByPeriodAndCategory(checkinDate, checkoutDate time.Time, category string) ([]*Room, error) {
	stmt, err := rd.db.Prepare(`select room_id, category, price, status from design_patterns.room where category = $1 and status = 'available' and 
	room_id not in (select room_id from design_patterns.booking where (checkin_date, checkout_date) overlaps ($2 ,$3) and status='confirmed')`)
	if err != nil {
		return nil, err
	}
	var rooms []*Room
	rows, err := stmt.Query(category, checkinDate, checkoutDate)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Room
		err = rows.Scan(&b.RoomId, &b.Category, &b.Price, &b.Status)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &b)
	}
	return rooms, nil
}

func (rb *RepositoryDatabase) ById(roomId int) (*Room, error) {
	stmt, err := rb.db.Prepare(`select room_id, category, price, status from design_patterns.room where room_id = $1`)
	if err != nil {
		return nil, err
	}
	var r Room
	rows, err := stmt.Query(roomId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&r.RoomId, &r.Category, &r.Price, &r.Status)
	}
	return &r, nil
}
