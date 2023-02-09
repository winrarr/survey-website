package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SurveyRoute(router *gin.Engine) {
	router.POST("/survey", controllers.CreateSurvey)
	router.GET("/surveys", controllers.GetAllSurveys)
}
