package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/app"
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

	// TODO: check for null parameters

	passcode := otp.GenerateOtp(otp.UnusedStrParam)

	if !otp.StorePasscodeOnRedis(passcode, user.UserName, otp.UnusedIntParam) {
		log.Printf("store otp on redis has been failed, username:%s\n", user.UserName)
	}

	// TODO: send passcode to email

	ctx.JSON(http.StatusOK, app.LoginResponse{
		Response: app.GeneralResponse{Status: http.StatusOK}, Code: passcode,
	})
}

func LoginConfirm(ctx *gin.Context) {

	var loginConfirmReq app.LoginConfirmRequest

	ctx.ShouldBindJSON(&loginConfirmReq)

	passcode := loginConfirmReq.Code
	storedPasscode := otp.GetPasscodeFromRedis(loginConfirmReq.Email)

	if passcode == "" || !otp.ValidateOtp(passcode, otp.UnusedStrParam) || passcode != storedPasscode {
		ctx.JSON(http.StatusBadRequest, app.LoginConfirmResponse{
			Response: app.GeneralResponse{Status: http.StatusBadRequest, Message: app.OtpIsNotValid},
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, app.LoginResponse{
		Response: app.GeneralResponse{Status: http.StatusOK},
	})
}
