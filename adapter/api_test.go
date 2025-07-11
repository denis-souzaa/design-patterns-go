package main

import (
	"denis-souzaa/design-patterns-go/adapter/internal/catalog"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSearchProductInCatalog(t *testing.T) {
	productId := 1
	response, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d", productId))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	output := catalog.Output{}
	json.Unmarshal(data, &output)
	assert.Equal(t, 1, output.ProductId)
	assert.Equal(t, "A", output.Description)
	assert.Equal(t, 100.0, output.Price)
}
