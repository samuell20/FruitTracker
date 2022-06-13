package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for Products creation.
func PutHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}
