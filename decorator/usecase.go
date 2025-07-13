package main

type Usecase interface {
	Execute(input Input) (*Output, error)
}
