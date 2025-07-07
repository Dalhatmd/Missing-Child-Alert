package controllers

import (
    "net/http"

    "github.com/dalhatmd/Missing-Child-Alert/db"
    "github.com/dalhatmd/Missing-Child-Alert/user"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// CreateUser handles the creation of a new user.
func CreateUser(c *gin.Context) {
    var userDTO user.UserDTO

    // Bind the JSON request body to the UserDTO struct.
    if err := c.ShouldBindJSON(&userDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Only require email and password
    if userDTO.Email == nil || *userDTO.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
        return
    }
    if userDTO.Password == nil || *userDTO.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
        return
    }

    userID := uuid.New().String()
    username := ""
    location := ""
    if userDTO.Username != nil {
        username = *userDTO.Username
    }
    if userDTO.Location != nil {
        location = *userDTO.Location
    }

    newUser, err := user.NewUser(userID, username, *userDTO.Email, *userDTO.Password, location)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := db.DB.Create(newUser)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, newUser)
}