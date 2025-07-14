package main

import (
	"denis-souzaa/design-patterns-go/mediator/average"
	"denis-souzaa/design-patterns-go/mediator/config"
	"denis-souzaa/design-patterns-go/mediator/grade"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestShouldCalculateAveradeGrade(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	src := rand.NewSource(time.Now().UnixNano())
	studentId := rand.New(src).Intn(1000)
	gradeRepo := grade.NewGradeRepositoryDatabase(db)
	avgRepo := average.NewAverageRepositoryDatabase(db)
	calcAvg := NewCalculate(gradeRepo, avgRepo)
	saveGrade := NewSave(gradeRepo, *calcAvg)
	input1 := grade.Input{
		StudentId: studentId,
		Exam:      "P1",
		Value:     10,
	}
	saveGrade.Execute(input1)

	input2 := grade.Input{
		StudentId: studentId,
		Exam:      "P2",
		Value:     9,
	}
	saveGrade.Execute(input2)

	input3 := grade.Input{
		StudentId: studentId,
		Exam:      "P3",
		Value:     8,
	}
	saveGrade.Execute(input3)
	avg := NewGetAverage(avgRepo)
	output := avg.Execute(studentId)
	assert.Equal(t, float64(9), output)
}

func TestShouldCalculateAveradeGradeWithMediator(t *testing.T) {
	db, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	src := rand.NewSource(time.Now().UnixNano())
	studentId := rand.New(src).Intn(1000)
	gradeRepo := grade.NewGradeRepositoryDatabase(db)
	avgRepo := average.NewAverageRepositoryDatabase(db)
	calcAvg := NewCalculate(gradeRepo, avgRepo)
	mediator := &Mediator{}
	mediator.Register("gradeSaved", func(data any) {
		calcAvg.Execute(int(data.(int)))
	})
	saveGrade := NewSaveMediator(gradeRepo, *mediator)
	input1 := grade.Input{
		StudentId: studentId,
		Exam:      "P1",
		Value:     10,
	}
	saveGrade.Execute(input1)

	input2 := grade.Input{
		StudentId: studentId,
		Exam:      "P2",
		Value:     9,
	}
	saveGrade.Execute(input2)

	input3 := grade.Input{
		StudentId: studentId,
		Exam:      "P3",
		Value:     8,
	}
	saveGrade.Execute(input3)
	avg := NewGetAverage(avgRepo)
	output := avg.Execute(studentId)
	assert.Equal(t, float64(9), output)
}
