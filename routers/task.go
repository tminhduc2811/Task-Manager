package routers

import (
	. "../controllers"
	"github.com/gin-gonic/gin"
)
func TaskRoutes(router *gin.RouterGroup, controller TaskController) *gin.RouterGroup {
	router.GET("/", controller.FindAll)
	router.GET("/:id", controller.FindById)
	//TODO: Add middleware for these routers below
	router.POST("/", controller.Create)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
	return router
}