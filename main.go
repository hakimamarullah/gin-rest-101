package main

import (
	"gin-rest/controllers/productcontrollers"
	"gin-rest/models"

	"github.com/gin-gonic/gin"
)

const (
	PRODUCTS = "/products"
	API_V1 = "/api/v1"
	PRODUCT_ID = "/product/:id"
	BATCH = "/batch"
)
func main() {
	router := gin.Default()
	models.ConnectDatabasePostgreSQL()

	v1 := router.Group(API_V1)
	{
		v1.GET(PRODUCTS, productcontrollers.GetAllProducts)
		v1.POST(PRODUCTS, productcontrollers.CreateProduct)
		v1.POST(BATCH + PRODUCTS, productcontrollers.BatchCreateProduct)
		v1.DELETE(PRODUCT_ID, productcontrollers.DeleteProductById)
		v1.PUT(PRODUCT_ID, productcontrollers.UpdateProductById)
		v1.GET(PRODUCT_ID, productcontrollers.GetProductById)
	}

	router.Run("localhost:8000")
}