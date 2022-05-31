package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/models"
)

func FindAllGoods(ctx *gin.Context) {
	var item []models.Good

	common.Db.Find(&item)

	ctx.JSON(http.StatusOK, gin.H{"message": item})
}
