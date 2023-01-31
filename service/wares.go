package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/model"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/utils"
	"log"
	"time"
)

type waresService struct{}

var WaresService waresService

// Add 添加商品
func (ws waresService) Add(query query.WaresAddQuery, c *gin.Context) error {
	var wares model.Wares

	loginUser := utils.GetLoginUser(c)
	wares.CreateAt = time.Now()
	wares.UpdateAt = time.Now()
	wares.Stock = query.Stock
	wares.UserId = loginUser.UId
	wares.Price = query.Price
	wares.Status = query.Status

	tx := model.DB.Model(&model.Wares{}).Create(&wares)
	return tx.Error
}

// Edit 编辑商品
func (ws waresService) Edit(query query.WaresEditQuery) error {
	var wares model.Wares
	wares.UpdateAt = time.Now()
	wares.Stock = query.Stock
	wares.Status = query.Status
	wares.Price = query.Price
	wares.Name = query.Name
	wares.WId = query.WId

	tx := model.DB.Model(&model.Wares{}).Where("w_id", wares.WId).Updates(&wares)
	return tx.Error
}

// Detail 商品详情
func (ws waresService) Detail(id uint64) (model.Wares, error) {
	var wares model.Wares
	tx := model.DB.Model(&model.Wares{}).Where("w_id", id).First(&wares)
	if tx.Error != nil {
		return wares, errors.New("未找到相关的商品信息")
	}
	return wares, nil
}

// Page 商品列表
func (ws waresService) Page(query query.WaresPageQuery) (utils.PageData, error) {
	var ret utils.PageData
	var rows []model.Wares

	tx := model.DB.Model(&model.Wares{})
	if query.Status != 0 {
		tx.Where("status", query.Status)
	}
	if query.Type != 0 {
		tx.Where("type", query.Type)
	}
	if query.BoothId != 0 {
		tx.Where("booth_id", query.BoothId)
	}
	tx.Offset(int((query.PageNo - 1) * query.PageSize)).Limit(int(query.PageSize))
	tx.Find(&rows)
	if tx.Error != nil {
		log.Printf("分页查询商品信息失败，err->%s\n", tx.Error)
		return ret, errors.New("查询商品失败")
	}
	return ret, nil
}
