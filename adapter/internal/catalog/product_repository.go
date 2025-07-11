package catalog

var products = []Product{
	{ProductId: 1, Description: "A", Price: 100.0},
	{ProductId: 2, Description: "B", Price: 200.0},
	{ProductId: 3, Description: "C", Price: 300.0},
}

type ProductRepository interface {
	GetById(p int) *Product
}

type ProductRepositoryMemory struct {
}

func (pm *ProductRepositoryMemory) GetById(productId int) *Product {
	for _, pr := range products {
		if pr.ProductId == productId {
			return &Product{ProductId: pr.ProductId, Description: pr.Description, Price: pr.Price}
		}
	}
	return nil
}
