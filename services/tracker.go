package services

import (
	"github.com/gin-gonic/gin"
)

func TrackerCreateNewUser(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "User Telah Ditambahkan",
		"response_data":    ""})
}
