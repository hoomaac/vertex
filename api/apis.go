package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/api/app"
	"github.com/hoomaac/vertex/api/auth"
)

func StartEngine(mode string) *gin.Engine {

	engine := gin.Default()

	gin.SetMode(mode)

	group := engine.Group("/api/v1")

	app.TellRoutes(group)

	auth.TellRoutes(group)

	return engine
}
