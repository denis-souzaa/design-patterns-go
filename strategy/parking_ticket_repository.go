package strategy

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ParkingTicketRepository interface {
	ByPlate(plate string) (*ParkingTicket, error)
	Save(p ParkingTicket) error
	Update(p ParkingTicket) error
}

type ParkingTicketRepositoryDatabase struct {
	db *sql.DB
}

func NewParkingTicketRepositoryDatabase(db *sql.DB) *ParkingTicketRepositoryDatabase {
	return &ParkingTicketRepositoryDatabase{db: db}
}

func (r *ParkingTicketRepositoryDatabase) ByPlate(plate string) (*ParkingTicket, error) {
	stmt, err := r.db.Prepare(`select * from design_patterns.parking_ticket where plate = $1`)
	if err != nil {
		return nil, err
	}
	var t ParkingTicket
	rows, err := stmt.Query(plate)
	if err != nil {
		return nil, err
	}
	var checkout sql.NullTime
	for rows.Next() {
		err = rows.Scan(&t.Plate, &t.CheckinDate, &checkout, &t.Fare, &t.Location)
		if checkout.Valid {
			t.CheckoutDate = checkout.Time
		}

		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func (r *ParkingTicketRepositoryDatabase) Save(p ParkingTicket) error {
	stmt, err := r.db.Prepare(`insert into design_patterns.parking_ticket(plate, checkin_date, location, fare) values($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(p.Plate, p.CheckinDate, p.Location, p.Fare)
	if err != nil {
		return nil
	}
	err = stmt.Close()
	if err != nil {
		return nil
	}
	return nil
}

func (r *ParkingTicketRepositoryDatabase) Update(p ParkingTicket) error {
	stmt, err := r.db.Prepare(`update design_patterns.parking_ticket set checkout_date = $1, fare = $2 where plate = $3`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.CheckoutDate, p.Fare, p.Plate)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
