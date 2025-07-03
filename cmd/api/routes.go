package main

import (
	"github.com/flvSantos15/go-category-msvc/cmd/api/controllers"
	"github.com/flvSantos15/go-category-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategoryController(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategoryController(ctx, inMemoryCategoryRepository)
	})
}
