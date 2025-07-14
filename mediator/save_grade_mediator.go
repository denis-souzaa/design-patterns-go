package main

import (
	"denis-souzaa/design-patterns-go/mediator/grade"
)

type SaveGradeMediator struct {
	GradeRepo grade.GradeRepository
	Mediator  Mediator
}

func NewSaveMediator(gradeRepo grade.GradeRepository, mediator Mediator) *SaveGradeMediator {
	return &SaveGradeMediator{GradeRepo: gradeRepo, Mediator: mediator}
}

func (sg *SaveGradeMediator) Execute(i grade.Input) {
	grade := &grade.Grade{StudentId: i.StudentId, Exam: i.Exam, Value: i.Value}
	sg.GradeRepo.Save(*grade)
	sg.Mediator.Notify("gradeSaved", i.StudentId)
}
