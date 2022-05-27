package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/models"
)

func DeleteGoods(ctx *gin.Context) {
	var item models.Goods

	if err := common.Db.Where("GID = ?", ctx.Param("GID")).First(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not found!"})
		return
	}

	common.Db.Delete(&item)

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
