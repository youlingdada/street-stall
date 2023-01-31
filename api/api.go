package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/utils"
)

func bind(c *gin.Context, obj interface{}) error {
	err := c.BindJSON(obj)
	if err != nil {
		log.Printf("参数绑定失败,err->%s\n", err)
		utils.DealResult(c, utils.Error("参数绑定错误"))
		return err
	}
	return nil
}
