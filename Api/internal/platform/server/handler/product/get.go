package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/get/product"
)

// CreateHandler returns an HTTP handler for Products creation.
func GetHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productService := service.(product.ProductQuery)
		idString, err := ctx.Params.Get("id")
		if err != true {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		id, parseError := strconv.ParseInt(idString, 10, 8)
		if parseError != nil {
			ctx.JSON(http.StatusBadRequest, "Not a number")
			return
		}
		product, productErr := productService.GetProductById(int(id), ctx)
		if productErr != nil {
			ctx.JSON(http.StatusBadRequest, "Error while fetching data")
			return
		}
		ctx.JSON(http.StatusOK, product)
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
