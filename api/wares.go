package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/service"
	"github.com/youlingdada/street-stall/utils"
	"log"
	"strconv"
)

type waresApi struct{}

var WaresApi waresApi

// Add 添加商品
// @Summary 添加商品
// @Schemes
// @Description 添加一个商品
// @Tags 商品模块
// @Param wares body query.WaresAddQuery true "商品信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /wares/add [post]
func (wApi waresApi) Add(c *gin.Context) {
	var wares query.WaresAddQuery
	err := bind(c, &wares)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
		return
	}
	if err := service.WaresService.Add(wares, c); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("添加商品成功"))
	}
}

// Edit 编辑商品
// @Summary 编辑商品
// @Schemes
// @Description 编辑一个商品
// @Tags 商品模块
// @Param wares body query.WaresEditQuery true "商品信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /wares/edit [post]
func (wApi waresApi) Edit(c *gin.Context) {
	var wares query.WaresEditQuery
	err := bind(c, &wares)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
		return
	}
	if err := service.WaresService.Edit(wares); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("更新商品成功"))
	}
}

// Detail 商品详情
// @Summary 商品详情
// @Schemes
// @Description 获取商品详细信息
// @Tags 商品模块
// @Accept json
// @Produce json
// @Param id query int64 true "商品id"
// @Success 200 {object} utils.Result
// @Router /wares/detail [get]
func (wApi waresApi) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if w, err := service.WaresService.Detail(uint64(id)); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("商品信息获取成功", w))
	}
}

// Page 商品列表
// @Summary 商品列表
// @Schemes
// @Description 分页获取商品信息
// @Tags 商品模块
// @Param wares body query.WaresPageQuery true "商品分页信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /wares/page [get]
func (wApi waresApi) Page(c *gin.Context) {
	var wares query.WaresPageQuery
	err := c.BindQuery(&wares)
	if err != nil {
		log.Printf("参数绑定失败，err->%s\n", err)
		utils.DealResult(c, utils.Error("获取商品列表失败"))
		return
	}
	if w, err := service.WaresService.Page(wares); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("商品列表获取成功", w))
	}
}
