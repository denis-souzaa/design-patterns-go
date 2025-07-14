package average

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type AverageRepository interface {
	Save(avg Average) error
	ByStudentId(studentId int) (*Average, error)
}

type AverageRepositoryDatabase struct {
	db *sql.DB
}

func NewAverageRepositoryDatabase(db *sql.DB) *AverageRepositoryDatabase {
	return &AverageRepositoryDatabase{db: db}
}

func (arg *AverageRepositoryDatabase) Save(avg Average) error {
	stmt, err := arg.db.Prepare(`delete from design_patterns.average where student_id = $1`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(avg.StudentId)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	stmt, err = arg.db.Prepare(`insert into design_patterns.average (student_id, value) values ($1, $2)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(avg.StudentId, avg.Value)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (arg *AverageRepositoryDatabase) ByStudentId(studentId int) (*Average, error) {
	stmt, err := arg.db.Prepare(`select * from design_patterns.average where student_id = $1`)
	if err != nil {
		return nil, err
	}
	var avg Average
	rows, err := stmt.Query(studentId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&avg.StudentId, &avg.Value)
		if err != nil {
			return nil, err
		}
	}

	return &avg, nil
}
