package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/delete/user"
	"github.com/samuell20/FruitTracker/kit/command"
)

type deleteRequest struct {
	Id int `json:"id" binding:"required"`
}

// CreateHandler returns an HTTP handler for Products creation.
func DeleteHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

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

		persistErr := commandBus.Dispatch(ctx, user.NewDeleteUserCommand(
			int(id),
		))
		if persistErr != nil {
			ctx.JSON(http.StatusBadRequest, "Error while persisting")
		}
		ctx.JSON(http.StatusOK, "OK")
	}
}
