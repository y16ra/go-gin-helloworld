package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/y16ra/go-gin-helloworld/handlers"
)

func SetupGreeting(r *gin.Engine) {
	r.GET("/greeting", handlers.Greeting())
	r.GET("/greeting/:name", handlers.GreetingTo())
	r.POST("/greeting", handlers.PostGreeting())
}
