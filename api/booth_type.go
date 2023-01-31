package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/service"
	"github.com/youlingdada/street-stall/utils"
	"strconv"
)

type boothTypeApi struct{}

var BoothTypeApi boothTypeApi

// Add 添加摊位类型
// @Summary 添加摊位类型
// @Schemes
// @Description 添加一个摊位类型信息
// @Tags 摊位类型模块
// @Param boothType body query.BoothTypeAddQuery true "摊位类型信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /boothType/add [post]
func (btApi boothTypeApi) Add(c *gin.Context) {
	var bt query.BoothTypeAddQuery
	err := bind(c, &bt)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
		return
	}
	if err = service.BoothTypeService.Add(bt); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("添加摊位类型成功"))
	}
}

// Edit 编辑摊位类型
// @Summary 编辑摊位类型
// @Schemes
// @Description 编辑一个摊位类型信息
// @Tags 摊位类型模块
// @Param boothType body query.BoothTypeEditQuery true "摊位类型信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /boothType/edit [post]
func (btApi boothTypeApi) Edit(c *gin.Context) {
	var bt query.BoothTypeEditQuery
	err := bind(c, &bt)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
		return
	}
	if err = service.BoothTypeService.Edit(bt); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("更新摊位类型成功"))
	}
}

// Detail 获取摊位详情
// @Summary 获取摊位详情
// @Schemes
// @Description 获取一个摊位类型信息
// @Tags 摊位类型模块
// @Param id query uint32 true "摊位类型id"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /boothType/detail [get]
func (btApi boothTypeApi) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	bt, err := service.BoothTypeService.Detail(uint32(id))
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("获取摊位类型成功", bt))
	}
}

// Page 获取摊位类型列表
// @Summary 获取摊位类型列表
// @Schemes
// @Description 获取摊位类型列表
// @Tags 摊位类型模块
// @Param pageNo query uint64 true "页号"
// @Param pageSize query uint64 true "每页的数据量"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /boothType/page [get]
func (btApi boothTypeApi) Page(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	bt, err := service.BoothTypeService.Page(uint64(pageNo), uint64(pageSize))
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("获取摊位类型列表成功", bt))
	}
}
