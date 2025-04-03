package services

import (
	"voip-manager-api/db"
	"voip-manager-api/models"
	"voip-manager-api/helpers"

	"net/http"
	"github.com/gin-gonic/gin"
)


func GetCallLogs(c *gin.Context) {
	var logs []models.CallLog
	db.DB.Preload("Voicemails").Find(&logs)
	c.JSON(http.StatusOK, logs)
}

func CreateCallLog(c *gin.Context) {
	var logEntry models.CallLog
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&logEntry)
	c.JSON(http.StatusCreated, logEntry)
}

func UpdateCallLog(c *gin.Context) {
	var logEntry models.CallLog
	id := c.Param("id")

	if err := db.DB.Preload("Voicemails").First(&logEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call log not found"})
		return
	}

	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&logEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update call log"})
		return
	}

	c.JSON(http.StatusOK, logEntry)
}

func DeleteCallLog(c *gin.Context) {
	id := c.Param("id")
	var logEntry models.CallLog

	if err := db.DB.First(&logEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call log not found"})
		return
	}

	if err := db.DB.Delete(&logEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete call log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Call log deleted successfully"})
}

func GetCallLogByID(c *gin.Context) {
	id := c.Param("id")
	var logEntry models.CallLog

	if err := db.DB.Preload("Voicemails").First(&logEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call log not found"})
		return
	}

	c.JSON(http.StatusOK, logEntry)
}

func SearchCallLogs(c *gin.Context) {
	var logs []models.CallLog

	query := helpers.BuildSearchCallLogQuery(c)

	if err := query.Preload("Voicemails").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search call logs"})
		return
	}

	if len(logs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No call logs found"})
		return
	}

	c.JSON(http.StatusOK, logs)
}