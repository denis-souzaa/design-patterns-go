package main

import (
	"denis-souzaa/design-patterns-go/mediator/average"
	"log"
)

type GetAverage struct {
	AvgRepo average.AverageRepository
}

func NewGetAverage(avgRepo average.AverageRepository) *GetAverage {
	return &GetAverage{AvgRepo: avgRepo}
}

func (gt *GetAverage) Execute(studendId int) float64 {
	average, err := gt.AvgRepo.ByStudentId(studendId)
	if err != nil {
		log.Fatal(err)
	}
	return average.Value
}
