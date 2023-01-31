/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-21 10:36:34
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-21 10:36:54
 * @FilePath: \street-stall\router\utils.go
 * @Description: 工具接口
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/api"
)

func initUtils(r *gin.Engine) {
	utils := r.Group("/utils")

	utils.GET("/getPublicRsaKey", api.GetPublicRsaKey)
	utils.GET("/encodeRsa", api.EncodeRsa)
	utils.GET("/decodeRsa", api.DecodeRsa)
}
