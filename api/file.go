package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/service"
	"github.com/youlingdada/street-stall/utils"
	"strconv"
)

type fileApi struct{}

var FileApi fileApi

// Add 添加文件
// @Summary 添加文件
// @Schemes
// @Description 添加一个文件
// @Tags 文件模块
// @Param file formData file true "文件"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /file/add [post]
func (fApi fileApi) Add(c *gin.Context) {
	id, err := service.FileService.Add(c)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("文件上传成功", id))
	}
}

// Del 删除文件
// @Summary 删除文件
// @Schemes
// @Description 根据文件id删除文件
// @Tags 文件模块
// @Param id query uint64 true "文件id"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /file/del [post]
func (fApi fileApi) Del(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := service.FileService.Del(uint64(id))
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("文件删除成功"))
	}
}

// Detail 文件详情
// @Summary 文件详情
// @Schemes
// @Description 根据文件id获取文件详情
// @Tags 文件模块
// @Param id query uint64 true "文件id"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /file/detail [get]
func (fApi fileApi) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fileVo, err := service.FileService.Detail(uint64(id))
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("文件信息获取成功", fileVo))
	}
}
