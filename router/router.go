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
		apiV1.DELETE("/sent-mails/:id", api.DeleteSentMail)
		apiV1.POST("/sent-mails/any/batch-deletions", api.DeleteSentMailsInBatch)

		apiV1.GET("/contacts", api.ListContacts)
		apiV1.POST("/contacts", api.AddContact)
		apiV1.DELETE("/contacts/:id", api.DeleteContact)
		apiV1.POST("/contacts/any/batch-deletions", api.DeleteContactsInBatch)

		apiV1.GET("/postmark-collections", api.ListPostmarkCollections)
		apiV1.POST("/postmark-collections", api.AddPostmarkCollection)
		apiV1.DELETE("/postmark-collections/:id", api.DeletePostmarkCollection)
		apiV1.POST("/postmark-collections/any/batch-deletions", api.DeletePostmarkCollectionsInBatch)
	}
	// swagger router and BasePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//docs.SwaggerInfo.BasePath = "/ai-task-dist/api/v1"
	docs.SwaggerInfo.BasePath = "/"
}
