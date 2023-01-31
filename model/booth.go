// Package model /*
package model

import "gorm.io/gorm"

type Booth struct {
	BId       uint64  `json:"bId" gorm:"column:b_id;primaryKey"` //摊位id
	Type      uint32  `json:"type" gorm:"column:type"`           // 摊位类型
	Address   string  `json:"address" gorm:"column:address"`     // 摊位地址
	Status    uint8   `json:"status" gorm:"column:status"`       // 摊位状态
	Name      string  `json:"name" gorm:"column:name"`           // 摊位名称
	UserId    uint64  `json:"userId" gorm:"column:user_id"`      // 用户id
	PosterId  uint64  `json:"posterId" gorm:"column:poster_id"`  // 摊位海报图片id
	Longitude float64 `json:"longitude" gorm:"column:longitude"` // 经度
	Latitude  float64 `json:"latitude" gorm:"column:latitude"`   // 维度
	Model
}

type boothModel struct{}

var BoothModel boothModel

// BoothDB 获取摊位DB
func (bm boothModel) BoothDB() *gorm.DB {
	return DB.Model(Booth{})
}
