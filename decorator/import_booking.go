package main

import (
	"log"
	"strings"
	"time"
)

type ImportBooking struct {
	usecase Usecase
}

func NewImportBooking(useCase Usecase) *ImportBooking {
	return &ImportBooking{usecase: useCase}
}

func (br *ImportBooking) Execute(i ImportInput) ([]string, error) {
	var checkinDate, checkoutDate time.Time
	var output []string
	lines := strings.Split(i.File, "\n")
	lines = lines[1:]

	for _, l := range lines {

		line := strings.Split(l, ";")
		checkinDate = parseDate(line[1])
		checkoutDate = parseDate(line[2])
		inputUsecase := Input{
			Email:        line[0],
			CheckinDate:  checkinDate,
			CheckoutDate: checkoutDate,
			Category:     line[3],
		}
		outputUsecase, err := br.usecase.Execute(inputUsecase)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, outputUsecase.Code)
	}

	return output, nil
}

type ImportInput struct {
	File string
}

func parseDate(date string) time.Time {
	value, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}

	return value
}
