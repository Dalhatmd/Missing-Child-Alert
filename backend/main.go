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

	router.POST("/users", controllers.CreateUser)

	router.Run(":8080")


}	