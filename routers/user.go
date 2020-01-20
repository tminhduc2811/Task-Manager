package routers

import (
	. "../controllers"
	. "../middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller UserController) *gin.RouterGroup {
	router.POST("/register", controller.Create)
	router.POST("login", AuthenticationMiddleware().LoginHandler)
	// Admin role for other methods
	auth := router.Group("/")
	auth.Use(AuthenticationMiddleware().MiddlewareFunc())
	{
		auth.GET("/", controller.FindAll)
		auth.GET("/:id", controller.FindById)
		auth.PUT("/:id", controller.Update)
		auth.DELETE("/:id", controller.Delete)
	}
	return router
}