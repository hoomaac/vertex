package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/models"
)

func AddGoods(ctx *gin.Context) {

	// Validate input
	var input models.Good
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	new_item := models.Good{Name: input.Name}

	common.Db.Create(&new_item)

	ctx.JSON(http.StatusOK, gin.H{"message": new_item})
}
