package booking

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository interface {
	Save(booking Booking) error
	Update(booking Booking) error
	ByCode(code string) (*Booking, error)
}

type RepositoryDatabase struct {
	db *sql.DB
}

func NewRepositoryDatabase(db *sql.DB) *RepositoryDatabase {
	return &RepositoryDatabase{
		db: db,
	}
}

func (rb *RepositoryDatabase) Save(b Booking) error {
	stmt, err := rb.db.Prepare(`
		insert into design_patterns.booking (code, room_id, email, checkin_date, checkout_date, duration, price, status) 
		values($1,$2,$3,$4,$5,$6,$7,$8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		b.Code,
		b.RoomId,
		b.Email,
		b.CheckinDate,
		b.CheckoutDate,
		b.Duration,
		b.Price,
		b.status,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (rb *RepositoryDatabase) Update(b Booking) error {
	stmt, err := rb.db.Prepare(`
		update design_patterns.booking set status=$1 where code=$2`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(
		b.status,
		b.Code,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (rb *RepositoryDatabase) ByCode(c string) (*Booking, error) {
	stmt, err := rb.db.Prepare(`select * from design_patterns.booking where code = $1`)
	if err != nil {
		return nil, err
	}
	var b Booking
	rows, err := stmt.Query(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&b.Code, &b.RoomId, &b.Email, &b.CheckinDate, &b.CheckoutDate, &b.Duration, &b.Price, &b.status)
	}
	return &b, nil
}
