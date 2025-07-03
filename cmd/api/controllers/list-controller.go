package controllers

import (
	"net/http"

	"github.com/flvSantos15/go-category-msvc/internal/repositories"
	usecases "github.com/flvSantos15/go-category-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategoryController(ctx *gin.Context, repository repositories.ICategoryRepository) {
	usecase := usecases.NewListCategoryUseCase(repository)

	categories, err := usecase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"categories": categories,
	})
}
