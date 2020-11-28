package routers

import (
	"api-starter/controllers"

	"api-starter/middlewares"

	"github.com/gin-gonic/gin"
)

func defineRoutes(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.Authentication())
	protectedRoutes.GET("/user/:email", authController.GetUser)
	protectedRoutes.PUT("/user", authController.UpdateUser)

}

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	defineRoutes(router)
	return router
}
