package main

import (
	"denis-souzaa/design-patterns-go/mediator/grade"
)

type SaveGrade struct {
	GradeRepo   grade.GradeRepository
	CalcAverage CalculateAverage
}

func NewSave(gradeRepo grade.GradeRepository, calcAverage CalculateAverage) *SaveGrade {
	return &SaveGrade{GradeRepo: gradeRepo, CalcAverage: calcAverage}
}

func (sg *SaveGrade) Execute(i grade.Input) {
	grade := &grade.Grade{StudentId: i.StudentId, Exam: i.Exam, Value: i.Value}
	sg.GradeRepo.Save(*grade)
	sg.CalcAverage.Execute(i.StudentId)
}
