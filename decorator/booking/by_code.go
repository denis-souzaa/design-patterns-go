package booking

type ByCode struct {
	Repo Repository
}

func (bc *ByCode) Execute(i Input) Output {
	b, err := bc.Repo.ByCode(i.Code)
	if err != nil {
		panic(err)
	}

	return Output{Duration: b.Duration, Price: b.Price, Code: b.Code}
}
