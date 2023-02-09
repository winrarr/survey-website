package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessions = map[string]time.Time{}

func NewSession() string {
	sessionToken := uuid.NewString()
	sessions[sessionToken] = time.Now()
	return sessionToken
}

func Authenticate(c *gin.Context) {
	sessionCookie, err := c.Request.Cookie("session_token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		println("hej1")
		return
	}

	val, ok := sessions[sessionCookie.Value]
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		println("hej2")
		return
	}

	expires := val.Add(2 * time.Minute)
	if expires.Before(time.Now()) {
		c.AbortWithStatus(440)
	}
}
