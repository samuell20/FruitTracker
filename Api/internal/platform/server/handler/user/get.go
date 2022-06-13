package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/operations/get/user"
)

type getRequest struct {
	ID int `json:"id" binding:"required"`
}

// CreateHandler returns an HTTP handler for Products creation.
func GetHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userService := service.(user.UserQuery)
		var req getRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		userService.GetUserById(req.ID, ctx)
		ctx.Status(http.StatusOK)
	}
}

func GetAllHandler(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userService := service.(user.UserQuery)

		users, err := userService.GetAllUsers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}
