package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/api"
)

func initUser(r *gin.Engine) {
	user := r.Group("/user")

	user.POST("/login", api.UserApi.UserLogin)
	user.POST("/register", api.UserApi.UserRegister)
	user.POST("/updateInfo", api.UserApi.UserUpdateInfo)
	user.GET("/sendRegCode", api.UserApi.UserSendRegEmailCode)
	user.GET("/getLoginUser", api.UserApi.GetLoginUser)
	user.POST("/updateAvatar", api.UserApi.UserUpdateAvatar)
	user.GET("/sendModifyCode", api.UserApi.UserSendModifyPwdEmail)
	user.POST("/updatePwd", api.UserApi.UserUpdatePwd)
	user.POST("/modifyUsername", api.UserApi.UserModifyUsername)
}
