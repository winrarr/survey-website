package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	router.GET("/users", controllers.GetAllUsers)
}
