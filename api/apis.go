package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/api/v1/app"
	"github.com/hoomaac/vertex/api/v1/auth"
	"github.com/hoomaac/vertex/api/v1/goods"
	"github.com/hoomaac/vertex/middleware/requestid"
)

func StartEngine(mode string) *gin.Engine {

	engine := gin.Default()

	engine.Use(requestid.GetRequestId())

	gin.SetMode(mode)

	group := engine.Group("/api/v1")

	app.TellRoutes(group)
	auth.TellRoutes(group)
	goods.TellRoutes(group)

	return engine
}
