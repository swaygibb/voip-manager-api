package routes

import (
	"voip-manager-api/services"

	"github.com/gin-gonic/gin"
)

func RegisterCallLogRoutes(r *gin.Engine) {
	callLogsGroup := r.Group("/call-logs")
	{
		callLogsGroup.GET("/", services.GetCallLogs)
		callLogsGroup.POST("/", services.CreateCallLog)
		callLogsGroup.PUT("/:id", services.UpdateCallLog)
		callLogsGroup.DELETE("/:id", services.DeleteCallLog)
		callLogsGroup.GET("/:id", services.GetCallLogByID)
		callLogsGroup.GET("/search", services.SearchCallLogs)
	}
}