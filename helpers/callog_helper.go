package helpers

import (
	"voip-manager-api/db"
	"voip-manager-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildSearchCallLogQuery(c *gin.Context) *gorm.DB {
	caller := c.Query("caller")
	receiver := c.Query("receiver")
	duration := c.Query("duration")
	timestamp := c.Query("timestamp")

	query := db.DB.Model(&models.CallLog{})

	if caller != "" {
		query = query.Where("caller = ?", caller)
	}
	if receiver != "" {
		query = query.Where("receiver = ?", receiver)
	}
	if duration != "" {
		query = query.Where("duration = ?", duration)
	}
	if timestamp != "" {
		query = query.Where("timestamp = ?", timestamp)
	}

	return query
}
