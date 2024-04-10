package main

import (
	"github.com/Mamush-Meshesha/backend/actions"
	"github.com/Mamush-Meshesha/backend/events"
	"github.com/Mamush-Meshesha/backend/profile"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/register", actions.Register)

	router.POST("/login", actions.Login)

	router.POST("/reset", actions.Reset)
	router.POST("/profile", profile.Profile)
	router.POST("/upload", actions.FileUpload)
	router.POST("/like_notify", events.Notify)

	router.Run(":8480")
}
