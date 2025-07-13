package booking

type Cancel struct {
	Repo Repository
}

func (c *Cancel) Execute(i Input) {
	b, err := c.Repo.ByCode(i.Code)
	if err != nil {
		panic(err)
	}
	b.Cancel()
	c.Repo.Update(*b)
}
