package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/vo"
)

func GetLoginUser(c *gin.Context) *vo.UserVo {
	user, ok := c.Get("loginUser")
	// 判断是否存在
	if !ok {
		return nil
	}
	if loginUser, ok := user.(vo.UserVo); ok {
		return &loginUser
	} else {
		return nil
	}
}
