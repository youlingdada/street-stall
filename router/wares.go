package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/api"
)

func initWares(r *gin.Engine) {
	wares := r.Group("/wares")

	wares.POST("/add", api.WaresApi.Add)
	wares.POST("/edit", api.WaresApi.Edit)
	wares.GET("/detail", api.WaresApi.Detail)
	wares.GET("/page", api.WaresApi.Page)
}
