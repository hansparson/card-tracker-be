package schema

import (
	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, payload Success) {
	ctx.JSON(payload.Status, payload)
}

func WriteJsonResponse_Failed(ctx *gin.Context, data Failed) {
	ctx.JSON(data.Status, data)
}
