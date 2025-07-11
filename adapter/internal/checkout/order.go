package checkout

type Order struct {
	Items []Item
}

func (o *Order) AddProduct(p ProductDTO, quantity int) {
	o.Items = append(o.Items, Item{ProductId: p.ProductId, Price: p.Price, Quantity: quantity})
}

func (o *Order) Total() float64 {
	var total float64
	for _, i := range o.Items {
		total += i.Total()
	}
	return total
}
