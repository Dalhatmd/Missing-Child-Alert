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

	userResp := user.UserResponse{
		ID:          newUser.ID,
		Username:    newUser.Username,
		Email:       newUser.Email,
		Location:    newUser.Location,
		PhoneNumber: newUser.PhoneNumber,
		Alerts:      newUser.Alerts,
	}

	c.JSON(http.StatusCreated, userResp)
}

// loginUser handles user login.
func LoginUser(c *gin.Context) {
	var loginDTO user.UserDTO

	// Bind the JSON request body to the UserDTO struct.
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email and password
	if loginDTO.Email == nil || *loginDTO.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	if loginDTO.Password == nil || *loginDTO.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	// Find user by email
	var existingUser user.User
	if err := db.DB.Where("email = ?", *loginDTO.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Check password
	if err := user.CheckPassword(existingUser.PasswordHash, *loginDTO.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userResp := user.UserResponse{
		ID:          existingUser.ID,
		Username:    existingUser.Username,
		Email:       existingUser.Email,
		Location:    existingUser.Location,
		PhoneNumber: existingUser.PhoneNumber,
		Alerts:      existingUser.Alerts,
	}

	c.JSON(http.StatusOK, userResp)
}
