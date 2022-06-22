package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/create/user"
	"github.com/samuell20/FruitTracker/kit/command"
)

type postRequest struct {
	Id        int    `json:"id"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	CompanyId int    `json:"company_id" `
	VatId     int    `json:"vat_id" `
}

// CreateHandler returns an HTTP handler for Products creation.
func PostHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req postRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err := commandBus.Dispatch(ctx, user.NewUserCommand(
			req.Id,
			req.Username,
			req.CompanyId,
			req.Email,
			req.Password,
			req.Role,
			req.VatId,
		))
		if err != nil {
			ctx.AbortWithError(http.StatusNotModified, err)
		}

		ctx.Status(http.StatusOK)
	}
}
