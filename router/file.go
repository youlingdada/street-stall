package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/api"
)

func initFile(r *gin.Engine) {
	file := r.Group("/file")
	file.POST("/add", api.FileApi.Add)
	file.POST("/del", api.FileApi.Del)
	file.GET("/detail", api.FileApi.Detail)
}
