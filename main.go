package main

import (
	"log"
	"test-talkjs/handler"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	router := gin.Default()

	loadEnv()

	router.GET("", handler.App)
	userRoute := router.Group("users")
	{
		userRoute.GET("", handler.Users)
		userRoute.GET("/:id", handler.GetUser)
		userRoute.GET("/:id/conversations", handler.UserConversations)
		userRoute.POST("/:id", handler.CreateUser)
	}

	userPresenceRoute := router.Group("presences")
	{
		userPresenceRoute.POST("", handler.Presences)
	}

	conversationsRoute := router.Group("conversations")
	{
		conversationsRoute.GET("/:id", handler.GetConverstaions)
		conversationsRoute.POST("/:id", handler.CreateConversations)
		conversationsRoute.DELETE("/:id", handler.DeleteConversations)
	}

	if err := router.Run(":10000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		log.Println("failed to load from .env")
	}
}
