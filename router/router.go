package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"my-postcrossing/api"
	"my-postcrossing/docs"
)

func RegisterRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/sent-mails", api.ListSentMails)
		apiV1.POST("/sent-mails", api.SendMail)
		apiV1.DELETE("/sent-mails/:id", api.DeleteSentMails)

		apiV1.POST("/contacts", api.AddContact)
	}
	// swagger router and BasePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//docs.SwaggerInfo.BasePath = "/ai-task-dist/api/v1"
	docs.SwaggerInfo.BasePath = "/"
}
