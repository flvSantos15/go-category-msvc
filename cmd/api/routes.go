package main

import (
	controllers "github.com/flvSantos15/go-category-msvc/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	
	categoryRoutes.POST("/", controllers.CreateCategoryController)
}