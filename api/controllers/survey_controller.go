package controllers

import (
	"api/configs"
	"api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var surveyCollection *mongo.Collection = configs.GetCollection(configs.DB, "surveys")

func CreateSurvey(c *gin.Context) {
	var survey models.Survey

	if err := c.BindJSON(&survey); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		println("hej1")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := surveyCollection.InsertOne(ctx, survey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		println("hej2")
		return
	}

	c.Status(http.StatusCreated)
}

func GetAllSurveys(c *gin.Context) {
	var users []models.Survey

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		println("hej1")
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var survey models.Survey
		if err = results.Decode(&survey); err != nil {
			c.Copy().AbortWithStatus(http.StatusInternalServerError)
			println("hej2")
			return
		}

		users = append(users, survey)
	}

	c.IndentedJSON(http.StatusOK, users)
}
