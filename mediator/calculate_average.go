package main

import (
	"denis-souzaa/design-patterns-go/mediator/average"
	"denis-souzaa/design-patterns-go/mediator/grade"
	"log"
)

type CalculateAverage struct {
	GradeRepo grade.GradeRepository
	AvgRepo   average.AverageRepository
}

func NewCalculate(gradeRepo grade.GradeRepository, avgRepo average.AverageRepository) *CalculateAverage {
	return &CalculateAverage{GradeRepo: gradeRepo, AvgRepo: avgRepo}
}

func (c *CalculateAverage) Execute(studentId int) {
	var total float64
	grades, err := c.GradeRepo.ListByStudentId(studentId)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range grades {
		total += g.Value
	}
	value := total / float64(len(grades))
	average := average.New(studentId, value)
	c.AvgRepo.Save(*average)
}
