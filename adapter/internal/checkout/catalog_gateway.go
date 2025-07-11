package checkout

import (
	"encoding/json"
	"fmt"
)

type CatalogGateway interface {
	GetProduct(productId int) ProductDTO
}

type ProductDTO struct {
	ProductId   int
	Description string
	Price       float64
}

type CatalogGatewayHttp struct {
	HttpClient HttpClient
}

func (cg *CatalogGatewayHttp) GetProduct(productId int) ProductDTO {
	response, err := cg.HttpClient.Get(fmt.Sprintf("http://localhost:3000/products/%d", productId))
	if err != nil {
		panic(err)
	}
	product := &ProductDTO{}
	json.Unmarshal(response, product)
	return *product
}
