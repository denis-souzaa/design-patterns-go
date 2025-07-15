package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Save(u User) error
	Update(u User) error
	Delete(email string) error
	List() ([]*User, error)
	ByEmail(email string) (*User, error)
}

type RepositoryDatabase struct {
	db *sql.DB
}

func NewRepositoryDatabase(db *sql.DB) *RepositoryDatabase {
	return &RepositoryDatabase{db: db}
}

func (rdb *RepositoryDatabase) Save(u User) error {
	stmt, err := rdb.db.Prepare(`insert into design_patterns.user (name, email, password, status) values ($1, $2,$3,$4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Name(), u.Email(), u.Password(), u.Status())
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RepositoryDatabase) Update(u User) error {
	stmt, err := rdb.db.Prepare(`update design_patterns.user set name = $1, email = $2, password = $3, status = $4 where email = $5`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Name(), u.Email(), u.Password(), u.Status(), u.Email())
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
func (rdb *RepositoryDatabase) Delete(email string) error {
	stmt, err := rdb.db.Prepare(`delete from design_patterns.user where email = $1`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RepositoryDatabase) List() ([]User, error) {
	stmt, err := rdb.db.Prepare(`select * from design_patterns.user`)
	if err != nil {
		return nil, err
	}
	var users []User
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name, mail, pass, status string
		err := rows.Scan(&name, &mail, &pass, &status)
		if err != nil {
			return nil, err
		}
		users = append(users, *new(name, mail, pass, status))
	}
	return users, nil
}

func (rdb *RepositoryDatabase) ByEmail(mail string) (*User, error) {
	stmt, err := rdb.db.Prepare(`select * from design_patterns.user where email = $1`)
	if err != nil {
		return nil, err
	}
	var u *User
	rows, err := stmt.Query(mail)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name, mail, pass, status string
		err := rows.Scan(&name, &mail, &pass, &status)
		if err != nil {
			return nil, err
		}
		u = new(name, mail, pass, status)
	}
	return u, nil
}
