package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/update/user"
	"github.com/samuell20/FruitTracker/kit/command"
)

type putRequest struct {
	Id        int    `json:"id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" default:""`
	Role      string `json:"role" binding:"required"`
	CompanyId int    `json:"company_id" default:"0"`
	VatId     int    `json:"vat_id" default:"0"`
}

// CreateHandler returns an HTTP handler for Products creation.
func PutHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req putRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		commandBus.Dispatch(ctx, user.NewUpdateUserCommand(
			req.Id,
			req.Username,
			req.Email,
			req.Password,
			req.Role,
			req.CompanyId,
			req.VatId,
		))
		ctx.Status(http.StatusOK)
	}
}
