package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/database"
)

func DeleteGoods(ctx *gin.Context) {
	var item models.Good

	if err := database.Db.Where("GID = ?", ctx.Param("GID")).First(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not found!"})
		return
	}

	database.Db.Delete(&item)

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
