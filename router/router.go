package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/mdw"
	"github.com/youlingdada/street-stall/utils"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.Use(mdw.Cors())

	router.Use(mdw.JwtMiddleware(accessUrls))

	router.StaticFS("/static", http.Dir(utils.StaticRoot))

	// 设置路由
	initUser(router)
	initUtils(router)
	initBooth(router)
	initFile(router)
	initWares(router)

	Router = router
}
