package grade

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type GradeRepository interface {
	Save(grade Grade) error
	ListByStudentId(studentId int) ([]*Grade, error)
}

type GradeRepositoryDatabase struct {
	db *sql.DB
}

func NewGradeRepositoryDatabase(db *sql.DB) *GradeRepositoryDatabase {
	return &GradeRepositoryDatabase{db: db}
}

func (grd *GradeRepositoryDatabase) Save(g Grade) error {
	stmt, err := grd.db.Prepare(`insert into design_patterns.grade (student_id, exam, value) values ($1,$2,$3)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		g.StudentId,
		g.Exam,
		g.Value,
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

func (grd *GradeRepositoryDatabase) ListByStudentId(studentId int) ([]*Grade, error) {
	stmt, err := grd.db.Prepare(`select student_id, exam, value from design_patterns.grade where student_id = $1`)
	if err != nil {
		return nil, err
	}

	var grades []*Grade
	rows, err := stmt.Query(studentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var g Grade
		err = rows.Scan(&g.StudentId, &g.Exam, &g.Value)
		if err != nil {
			return nil, err
		}
		grades = append(grades, &g)
	}

	return grades, nil
}
