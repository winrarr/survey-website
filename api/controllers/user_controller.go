package controllers

import (
	"api/auth"
	"api/configs"
	"api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		println("hej1")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		println("hej3")
		return
	}

	c.Status(http.StatusCreated)
}

func Login(c *gin.Context) {
	var login models.Login

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		println("hej1")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": login.Username}).Decode(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		println("hej2")
		return
	}

	c.SetCookie("session_token", auth.NewSession(), 2, "/", "localhost", false, false)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

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
		var user models.User
		if err = results.Decode(&user); err != nil {
			c.Copy().AbortWithStatus(http.StatusInternalServerError)
			println("hej2")
			return
		}

		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}
