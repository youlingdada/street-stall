package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/service"
	"github.com/youlingdada/street-stall/utils"
)

type boothApi struct{}

var BoothApi boothApi

// Add 摊位添加
// @Summary 添加摊位
// @Schemes
// @Description 添加一个摊位的信息
// @Tags 摊位模块
// @Param booth body query.BoothAddQuery true "摊位信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /booth/add [post]
func (ba boothApi) Add(c *gin.Context) {
	var query query.BoothAddQuery
	err := bind(c, &query)
	if err != nil {
		return
	}
	err = service.BoothService.Add(query)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("摊位添加成功"))
	}
}

// Edit 摊位编辑
// @Summary 编辑摊位
// @Schemes
// @Description 编辑一个摊位的信息
// @Tags 摊位模块
// @Param booth body query.EditBoothQuery true "摊位信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /booth/edit [post]
func (ba boothApi) Edit(c *gin.Context) {
	var query query.EditBoothQuery
	err := bind(c, &query)
	if err != nil {
		return
	}
	err = service.BoothService.Edit(query)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("摊位修改成功"))
	}
}

// Detail 摊位信息
// @Summary 摊位详情
// @Schemes
// @Description 获取一个摊位的详情
// @Tags 摊位模块
// @Param id query int64 true "摊位id"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /booth/detail [get]
func (ba boothApi) Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Printf("获取摊位id失败,err->%s", err)
		utils.DealResult(c, utils.Error("摊位id未填写"))
	}
	b, err := service.BoothService.Detail(uint64(id))
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("获取摊位信息成功", b))
	}
}

// Page 摊位列表
// @Summary 摊位列表
// @Schemes
// @Description 获取摊位列表
// @Tags 摊位模块
// @Param query query query.PageBoothQuery true "摊位分页信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /booth/page [get]
func (ba boothApi) Page(c *gin.Context) {
	var query query.PageBoothQuery
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	query.PageNo = uint64(pageNo)
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	query.PageSize = uint64(pageSize)
	bType, _ := strconv.Atoi(c.Query("type"))
	query.Type = uint8(bType)
	status, _ := strconv.Atoi(c.Query("status"))
	query.Status = uint8(status)
	userId, _ := strconv.Atoi(c.Query("userId"))
	query.UserId = uint64(userId)

	/*	err := c.BindQuery(&query)
		if err != nil {
			log.Printf("参数绑定失败,err->%s", err)
			utils.DealResult(c, utils.Error("参数绑定失败"))
		}*/
	pd, err := service.BoothService.Page(query)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("获取摊位列表成功", pd))
	}
}
