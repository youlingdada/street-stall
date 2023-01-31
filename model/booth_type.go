package model

import (
	"github.com/youlingdada/street-stall/utils"
)

// BoothType 摊位类型
type BoothType struct {
	TId         uint32 `json:"tId" gorm:"column:t_id;primaryKey"`     // 类型id
	Name        string `json:"name" gorm:"column:name"`               // 类型名称
	Description string `json:"description" gorm:"column:description"` // 类型描述
	IconId      uint64 `json:"iconId" gorm:"column:icon_id"`          // 图标文件id
	Status      uint8  `json:"status" gorm:"column:status"`           // 类型状态 1 正常 2 下架
	Model
}
type boothTypeModel struct{}

var BoothTypeModel boothTypeModel

// Add 添加一个摊位类型
func (btm boothTypeModel) Add(boothType BoothType) error {
	return DB.Model(&BoothType{}).Create(&boothType).Error
}

// Edit 编辑一个摊位类型
func (btm boothTypeModel) Edit(boothType BoothType) error {
	return DB.Model(&BoothType{}).Where("t_id", boothType.TId).Updates(&boothType).Error
}

// Detail 获取一个摊位详情
func (btm boothTypeModel) Detail(id uint32) (BoothType, error) {
	var boothType BoothType
	tx := DB.Model(&BoothType{}).Where("t_id", id).Find(&boothType)
	return boothType, tx.Error
}

// Page 分页获取摊位类型
func (btm boothTypeModel) Page(order, size uint64) (utils.PageData, error) {
	var pageData utils.PageData
	var rows []BoothType
	tx := DB.Model(&BoothType{}).Where("status", 1)
	// 查询总数
	var count int64
	tx.Count(&count)
	pageData.TotalRow = uint64(count)
	if count > 0 {
		pageData.TotalPage = uint64(count)/size + func() uint64 {
			if uint64(count)%size == 0 {
				return 0
			}
			return 1
		}()
	}
	tx.Offset(int(order)).Limit(int(size)).Find(&rows)
	pageData.Rows = rows
	return pageData, tx.Error
}
