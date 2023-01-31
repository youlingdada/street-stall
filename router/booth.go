package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/api"
)

func initBooth(r *gin.Engine) {
	booth := r.Group("/booth")

	bApi := api.BoothApi
	booth.POST("/add", bApi.Add)
	booth.POST("/edit", bApi.Edit)
	booth.GET("/detail", bApi.Detail)
	booth.GET("/page", bApi.Page)
}
