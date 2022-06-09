package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/app"
)

func Register(ctx *gin.Context) {

	var registerReq app.RegisterRequest

	err := ctx.ShouldBindJSON(&registerReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, app.RegisterResponse{
			Reponse: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserRegisteredNotOk},
		})
		return
	}

	newUser := models.CreateUser(&registerReq)

	if newUser == nil {
		ctx.JSON(http.StatusBadRequest, app.RegisterResponse{
			Reponse: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserRegisteredNotOk},
		})
		return
	}

	// TODO: send verification email
	ctx.JSON(http.StatusOK, app.RegisterResponse{
		Reponse: app.GeneralResponse{Status: http.StatusOK, Message: app.UserRegisteredOk},
	})
}
