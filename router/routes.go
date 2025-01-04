package router

import (
	"github.com/ItKarma/learning_go_api/controller"
	"github.com/gin-gonic/gin"
)

func initalizeRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/product", controller.ShowPrudctController)

		apiGroup.POST("/products", controller.CreatePrudctController)

		apiGroup.DELETE("/products", controller.DeletePrudctController)

		apiGroup.PUT("/products", controller.UpdatePrudctController)

		apiGroup.GET("/products", controller.ListPrudctController)
	}
}
