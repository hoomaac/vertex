package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/pkg/database"
	"github.com/hoomaac/vertex/models"
)

func UpdateGoods(ctx *gin.Context) {
	var item models.Good

	if err := database.Db.Where("GID = ?", ctx.Param("GID")).First(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not found!"})
		return
	}

	// Validate input
	var updated_item models.Good
	if err := ctx.BindJSON(&updated_item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.Db.Model(&item).Updates(updated_item)

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
