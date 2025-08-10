package controllers

import (
	"net/http"

	"github.com/dalhatmd/Missing-Child-Alert/alert"
	"github.com/dalhatmd/Missing-Child-Alert/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAlert handles the creation of a new alert.
func CreateAlert(c *gin.Context) {
	var alertDTO alert.AlertDTO

	if err := c.ShouldBindJSON(&alertDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alertID := uuid.New().String()
	newAlert, err := alert.NewAlert(alertID, alertDTO.ChildName, alertDTO.Age, alertDTO.Gender, alertDTO.LastSeenLocation, alertDTO.Description, alertDTO.PhotoUrl, alertDTO.ReporterContact, alertDTO.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.DB.Create(newAlert); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newAlert)
}

// UpdateAlert handles updating an existing alert by ID.
func UpdateAlert(c *gin.Context) {
	id := c.Param("id")

	var alertDTO alert.AlertDTO // Use the same DTO for updates
	if err := c.ShouldBindJSON(&alertDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAlert alert.Alert
	if err := db.DB.First(&existingAlert, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	// Update fields if provided in DTO
	if alertDTO.ChildName != "" {
		existingAlert.ChildName = alertDTO.ChildName
	}
	if alertDTO.Age != 0 {
		existingAlert.Age = alertDTO.Age
	}
	if alertDTO.Gender != "" {
		existingAlert.Gender = alertDTO.Gender
	}
	if alertDTO.LastSeenLocation != "" {
		existingAlert.LastSeenLocation = alertDTO.LastSeenLocation
	}
	if alertDTO.Description != "" {
		existingAlert.Description = alertDTO.Description
	}
	if alertDTO.PhotoUrl != "" {
		existingAlert.PhotoUrl = alertDTO.PhotoUrl
	}
	if alertDTO.ReporterContact != "" {
		existingAlert.ReporterContact = alertDTO.ReporterContact
	}
	if alertDTO.UserId != "" {
		existingAlert.UserId = alertDTO.UserId
	}

	if result := db.DB.Save(&existingAlert); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAlert)
}

// DeleteAlert handles deleting an alert by ID.
func DeleteAlert(c *gin.Context) {
	id := c.Param("id")

	var existingAlert alert.Alert
	if err := db.DB.First(&existingAlert, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	if result := db.DB.Delete(&existingAlert); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}

// ResolveAlert sets an alert's status to resolved.
func ResolveAlert(c *gin.Context) {
	id := c.Param("id")

	var existingAlert alert.Alert
	if err := db.DB.First(&existingAlert, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	// Set the status to resolved (adjust this value to match your AlertStatus type)
	existingAlert.Status = "resolved" // or alert.AlertStatusResolved if using a custom type

	if result := db.DB.Save(&existingAlert); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAlert)
}
