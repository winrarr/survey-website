package controllers

import (
	"api/auth"
	"api/configs"
	"api/models"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userCollection.Drop(ctx)
}

func CreateUser(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	// user := models.User{
	// 	Username: c.PostForm("username"),
	// 	Password: c.PostForm("password"),
	// 	Name:     c.PostForm("name"),
	// }

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func Login(c *gin.Context) {
	var login models.Login

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": login.Username}).Decode(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
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
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var user models.User
		if err = results.Decode(&user); err != nil {
			c.Copy().AbortWithStatus(http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}
