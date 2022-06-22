package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/update/product"
	"github.com/samuell20/FruitTracker/kit/command"
)

type putRequest struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Unit        string  `json:"unit" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	TypeId      int     `json:"typeId" default:"0"`
	DiscountId  int     `json:"discountId" default:"0"`
	TaxId       int     `json:"taxId" default:"0"`
}

// CreateHandler returns an HTTP handler for Products creation.
func PutHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req putRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		commandBus.Dispatch(ctx, product.NewUpdateProductCommand(
			req.Id, req.Name, req.Description, req.Unit, req.Price, req.TypeId, req.DiscountId, req.TaxId,
		))
		ctx.JSON(http.StatusOK, "Ok")
	}
}
