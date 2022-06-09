package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/database"
)

func FindAllGoods(ctx *gin.Context) {
	var item []models.Good

	database.Db.Find(&item)

	ctx.JSON(http.StatusOK, gin.H{"message": item})
}
