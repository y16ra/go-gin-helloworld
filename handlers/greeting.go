package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greeting() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	}
}

func GreetingTo() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		msg := &GreetingMessage{
			Message: "Hello, " + name + "!",
		}
		c.JSON(http.StatusOK, msg)
	}
}

func PostGreeting() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &struct {
			Name string
		}{}
		c.ShouldBindJSON(req)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + req.Name,
		})
	}
}

type GreetingMessage struct {
	Message string `json:"message"`
}
