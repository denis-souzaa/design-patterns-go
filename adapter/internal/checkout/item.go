package checkout

type Item struct {
	ProductId int
	Quantity  int
	Price     float64
}

func (i *Item) Total() float64 {
	return i.Price * float64(i.Quantity)
}
