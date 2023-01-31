/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-07 09:15:58
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-22 15:28:45
 * @FilePath: \street-stall\main.go
 * @Description:
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */
package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/youlingdada/street-stall/config"
	docs "github.com/youlingdada/street-stall/docs"
	"log"

	"github.com/youlingdada/street-stall/router"
)

func main() {

	// 读取配置，启动程序
	app := config.GlobalConfig
	addr := fmt.Sprintf("%s:%d", app.Server.Host, app.Server.Port)

	docs.SwaggerInfo.BasePath = "/"
	swagger := ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.Router.GET("/swagger/*any", swagger)
	err := router.Router.Run(addr)
	if err != nil {
		log.Printf("Route binding failed,err: %s\n", err)
		log.Panic("Application stop")
	}

}
