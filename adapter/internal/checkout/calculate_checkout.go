package checkout

type CalculateCheckout struct {
	CatalogGateway CatalogGateway
}

type Input struct {
	Items []ItemInput
}

type ItemInput struct {
	ProductId int
	Quantity  int
}

func (cc *CalculateCheckout) Execute(i Input) float64 {
	order := &Order{}
	for _, i := range i.Items {
		product := cc.CatalogGateway.GetProduct(i.ProductId)
		order.AddProduct(product, i.Quantity)
	}
	total := order.Total()
	return total
}
