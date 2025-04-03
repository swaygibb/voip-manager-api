package services

import (
	"voip-manager-api/db"
	"voip-manager-api/models"
	"voip-manager-api/helpers"

	"net/http"
	"github.com/gin-gonic/gin"
)

func GetVoicemails(c *gin.Context) {
	var voicemails []models.Voicemail

	err := db.DB.Find(&voicemails).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch voicemails"})
		return
	}

	c.JSON(http.StatusOK, voicemails)
}

func CreateVoicemail(c *gin.Context) {
	var voicemail models.Voicemail

	if err := c.ShouldBindJSON(&voicemail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.DB.Create(&voicemail).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create voicemail"})
		return
	}

	c.JSON(http.StatusCreated, voicemail)
}

func UpdateVoicemail(c *gin.Context) {
	var voicemail models.Voicemail
	id := c.Param("id")

	if err := db.DB.First(&voicemail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voicemail not found"})
		return
	}

	if err := c.ShouldBindJSON(&voicemail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.DB.Save(&voicemail).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update voicemail"})
		return
	}

	c.JSON(http.StatusOK, voicemail)
}

func DeleteVoicemail(c *gin.Context) {
	id := c.Param("id")
	var voicemail models.Voicemail

	if err := db.DB.First(&voicemail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voicemail not found"})
		return
	}

	if err := db.DB.Delete(&voicemail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete voicemail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voicemail deleted successfully"})
}

func GetVoicemail(c *gin.Context) {
	id := c.Param("id")
	var voicemail models.Voicemail

	if err := db.DB.First(&voicemail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voicemail not found"})
		return
	}

	c.JSON(http.StatusOK, voicemail)
}

func SearchVoicemails(c *gin.Context) {
	var voicemails []models.Voicemail

	query := helpers.BuildSearchVoicemailQuery(c)

	if err := query.Find(&voicemails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search voicemails"})
		return
	}

	if len(voicemails) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No voicemails found"})
		return
	}

	c.JSON(http.StatusOK, voicemails)
}