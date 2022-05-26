package auth

import (

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/common/vtypes"
	"github.com/hoomaac/vertex/models"
)

func TellRoutes(group *gin.RouterGroup) {

	group.Group("/auth").GET("/register", registerUser)
	group.Group("/auth").POST("/login", loginUser)
}

func registerUser(ctx *gin.Context) {

	user := models.User{}

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(vtypes.BadRequest, vtypes.AuthResponse{Status: vtypes.BadRequest, Data: err.Error()})
		return
	}

	code, resp := models.CreateUser(&user)

	ctx.JSON(code, resp)
}

func loginUser(ctx *gin.Context) {

}
