package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCalculateCheckout(t *testing.T) {
	input := Input{
		Items: []ItemInput{
			{ProductId: 1, Quantity: 1},
			{ProductId: 2, Quantity: 2},
			{ProductId: 3, Quantity: 3},
		},
	}

	httpClient := &HttpNativeAdapter{}
	catalogGateway := CatalogGatewayHttp{HttpClient: httpClient}
	calcCheckout := &CalculateCheckout{CatalogGateway: &catalogGateway}
	output := calcCheckout.Execute(input)
	assert.Equal(t, 1400.0, output)
}
