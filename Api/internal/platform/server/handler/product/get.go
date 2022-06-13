package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/get/product"
)

type getRequest struct {
	ID string `json:"id" binding:"required"`
}

// CreateHandler returns an HTTP handler for Products creation.
func GetHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func GetAllHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productService := service.(product.ProductQuery)

		products, err := productService.GetAllProducts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}
