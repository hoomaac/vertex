package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/controllers/v1/goods"
)

func TellRoutes(group *gin.RouterGroup) {

	group.POST("/goods", goods.AddGoods)
	group.GET("/goods/:GID", goods.FindGoods)
	group.GET("/goods", goods.FindAllGoods)
	group.PATCH("/goods/:GID", goods.UpdateGoods)
	group.DELETE("/goods/:GID", goods.DeleteGoods)
}
