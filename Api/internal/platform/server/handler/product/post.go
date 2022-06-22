package product

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/internal/operations/create/product"
	"github.com/samuell20/FruitTracker/kit/command"
)

type createRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Unit        string  `json:"unit" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	TypeId      int     `json:"type_id" default:"0"`
	DiscountId  int     `json:"discount_id" default:"0"`
	TaxId       int     `json:"tax_id" default:"0"`
}

// CreateHandler returns an HTTP handler for Products creation.
func PostHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, product.NewProductCommand(
			req.Name,
			req.Description,
			req.Unit,
			req.Price,
			req.TypeId,
			req.DiscountId,
			req.TaxId,
		))

		if err != nil {
			switch {
			case errors.Is(err, model.ErrInvalidProductId),
				errors.Is(err, model.ErrEmptyProductName):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, "Todo bien")
	}
}
