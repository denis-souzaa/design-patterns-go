package main

import (
	"denis-souzaa/design-patterns-go/adapter/internal/catalog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/products/:productId", func(c *gin.Context) {
		param := c.Param("productId")
		productId, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		productRepo := &catalog.ProductRepositoryMemory{}
		productGet := &catalog.ProductGet{ProductRepo: productRepo}
		output := productGet.Execute(int(productId))
		c.JSONP(200, output)
	})
	router.Run(":3000")
}
