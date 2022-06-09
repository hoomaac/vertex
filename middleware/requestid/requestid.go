package requestid

import (
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetRequestId() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		request_id := uuid.NewV4().String()

		log.Printf("Request Id: %s\n", request_id)

		ctx.Writer.Header().Set("X-Request-Id", request_id)

		ctx.Next()
	}
}
