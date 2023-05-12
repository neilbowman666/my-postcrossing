package router

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api"
)

func RegisterRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/sent-mails", api.SendMail)
	}
}
