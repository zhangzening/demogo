package router

import (
	"demogo/src/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init(engine *gin.Engine) {
	// use ginSwagger middleware to
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")

	api.POST("/animal", handler.Create)
	api.GET("/animals", handler.List)

	api.POST("/upload", handler.UploadFile)
	api.POST("/multi/upload", handler.MultiFileUpload)

}
