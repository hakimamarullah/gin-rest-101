package productcontrollers

import (
	"gin-rest/models"
	"gin-rest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, utils.Response("Success", map[string]interface{}{"products": products, "counts": len(products)}))

}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response("Invalid data", nil))
		return
	}

	models.DB.Create(&newProduct)
	c.JSON(http.StatusCreated, utils.Response("Product created", newProduct))
}

func BatchCreateProduct(c *gin.Context) {
	var newProducts []models.Product

	if err := c.ShouldBindJSON(&newProducts); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response("Invalid data", nil))
		return
	}

	models.DB.CreateInBatches(&newProducts, len(newProducts))
	c.JSON(http.StatusCreated, utils.Response("Products created", map[string]interface{}{"counts": len(newProducts)}))

}

func DeleteProductById(c *gin.Context) {
		if id, ok := c.Params.Get("id"); ok {
			result := models.DB.Delete(&models.Product{}, id)

			if result.RowsAffected == 0 {
				c.JSON(http.StatusNotFound, utils.Response("Product Not Found", nil))
				return
			}

			c.JSON(http.StatusOK, utils.Response("Product successfully deleted", nil))
			return
		}else{
			c.JSON(http.StatusBadRequest, utils.Response("Bad Request, Make sure id is present as parameter", nil))
		}

}

func UpdateProductById(c *gin.Context) {
	productDTO := struct {
		ProductName string `json:"product_name"`
		Description string `json:"description"`
	}{}

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response("Invalid data", nil))
		return
	}
}

func GetProductById(c *gin.Context) {

}