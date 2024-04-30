package handler

import (
	"net/http"

	"github.com/MasterTeus/go-works.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendRequestError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	oppening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	if err := db.Create(&oppening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendRequestError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendRequestSuccess(ctx, oppening)
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET All Opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
