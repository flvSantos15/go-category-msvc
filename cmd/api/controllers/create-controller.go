package controllers

import (
	"net/http"

	usecases "github.com/flvSantos15/go-category-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategoryController(ctx *gin.Context) {
	var body createCategoryInput
	
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		return
	}
	
	usecase := usecases.NewCreateCategoryUseCase()
	
	err := usecase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}