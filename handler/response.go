package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendRequestError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"messsage": message,
		"code":     statusCode,
	})
}

func sendRequestSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Succefuly",
		"data":    data,
	})
}
