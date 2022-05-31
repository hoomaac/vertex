package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/models"
)

func FindGoods(ctx *gin.Context) {
	var item models.Good

	if err := common.Db.Where("GID = ?", ctx.Param("GID")).First(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": item})
}
