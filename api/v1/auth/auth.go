package auth

import (
	"github.com/gin-gonic/gin"
	user "github.com/hoomaac/vertex/controllers/v1/users"
)

func TellRoutes(group *gin.RouterGroup) {

	group.Group("/auth").GET("/register", user.Register)
	group.Group("/auth").POST("/login", user.Login)
}
