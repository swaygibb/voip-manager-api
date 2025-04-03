package routes

import (
	"voip-manager-api/services"

	"github.com/gin-gonic/gin"
)

func RegisterVoicemailRoutes(r *gin.Engine) {
	voicemailGroup := r.Group("/voicemails")
	{
		voicemailGroup.GET("/", services.GetVoicemails)
		voicemailGroup.POST("/", services.CreateVoicemail)
		voicemailGroup.PUT("/:id", services.UpdateVoicemail)
		voicemailGroup.DELETE("/:id", services.DeleteVoicemail)
		voicemailGroup.GET("/:id", services.GetVoicemail)
		voicemailGroup.GET("/search", services.SearchVoicemails)
	}
}