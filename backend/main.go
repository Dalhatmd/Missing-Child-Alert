package main

import (
	//	"fmt"
	//	"time"
	//	"github.com/dalhatmd/Missing-Child-Alert/alert"
	"github.com/dalhatmd/Missing-Child-Alert/db"
	//	"github.com/dalhatmd/Missing-Child-Alert/user"
	//	"github.com/google/uuid"
	"github.com/dalhatmd/Missing-Child-Alert/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDatabase()

	router := gin.Default()

	// Serve static files from the frontend directory
	router.Static("/static", "../frontend")

	// Serve index.html at the root path
	router.GET("/", func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	// Serve auth.html for login/signup
	router.GET("/auth", func(c *gin.Context) {
		c.File("../frontend/auth.html")
	})

	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)
	router.Run(":8081")

}
