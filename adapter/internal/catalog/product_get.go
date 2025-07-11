package catalog

type ProductGet struct {
	ProductRepo ProductRepository
}

type Output struct {
	ProductId   int
	Description string
	Price       float64
}

func (p *ProductGet) Execute(productId int) Output {
	pr := p.ProductRepo.GetById(productId)
	return Output{ProductId: pr.ProductId, Description: pr.Description, Price: pr.Price}
}
