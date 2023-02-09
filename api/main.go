package main

import (
	"api/configs"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.Static("/static", "./static")
	routes.UserRoute(router)
	router.Use()
	routes.SurveyRoute(router)

	router.Run("localhost:8080")
}
