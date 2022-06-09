package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/app"
	"github.com/hoomaac/vertex/pkg/jwt"
	"github.com/hoomaac/vertex/pkg/otp"
)

func Login(ctx *gin.Context) {

	var loginReq app.LoginRequest

	ctx.ShouldBindJSON(&loginReq)

	user := models.FindUserByEmail(loginReq.Email)

	if user == nil {
		ctx.JSON(http.StatusBadRequest, app.LoginResponse{
			Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserNotFound},
		})
		return
	}

	passcode := otp.GenerateOtp(user.Email)

	if passcode == "" {
		ctx.JSON(http.StatusInternalServerError, app.LoginResponse{
			Response: app.GeneralResponse{Status: http.StatusInternalServerError, Message: app.InternalGenericError},
		})
		return
	}

	// TODO: send passcode to email

	ctx.JSON(http.StatusOK, app.LoginResponse{
		Response: app.GeneralResponse{Status: http.StatusOK}, Code: passcode,
	})
}

func LoginConfirm(ctx *gin.Context) {

	var loginConfirmReq app.LoginConfirmRequest

	ctx.ShouldBindJSON(&loginConfirmReq)

	user := models.FindUserByEmail(loginConfirmReq.Email)

	// If user is already veritifed, do not continue the login confirm process
	if user.Verified {
		ctx.JSON(http.StatusBadRequest, app.LoginConfirmResponse{
			Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserAlreadyConfirm},
		})
		return
	}

	if loginConfirmReq.Code == "" || loginConfirmReq.Email == "" || !otp.ValidateOtp(loginConfirmReq.Code, loginConfirmReq.Email) {
		ctx.JSON(http.StatusBadRequest, app.LoginConfirmResponse{
			Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.OtpIsNotValid},
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusBadRequest, app.LoginConfirmResponse{
			Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserNotFound},
		})
		return
	}

	// Update the user
	user.Verified = true

	if !models.UpdateUser(user) {
		if user == nil {
			ctx.JSON(http.StatusBadRequest, app.LoginConfirmResponse{
				Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.UserConfirmFailed},
			})
			return
		}
	}

	token := jwt.GenerateJwt(user.UserName, user.Email)

	ctx.JSON(http.StatusOK, app.LoginConfirmResponse{
		Response: app.GeneralResponse{Status: http.StatusOK, Message: app.UserConfirmSuccess}, Token: token,
	})
}
