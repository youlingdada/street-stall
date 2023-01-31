// Package model /*
package model

type Wares struct {
	WId         uint64  `json:"wId" gorm:"column:w_id;primaryKey"`     // 商品id
	Name        string  `json:"name" gorm:"column:name"`               // 名称
	Description string  `json:"description" gorm:"column:description"` // 描述
	Price       float64 `json:"price" gorm:"price"`                    // 价格
	Stock       uint32  `json:"stock" gorm:"stock"`                    // 库存
	Views       uint64  `json:"views" gorm:"views"`                    // 浏览量
	UserId      uint64  `json:"userId" gorm:"user_id"`                 // 用户id
	Status      uint8   `json:"status" gorm:"status"`                  // 状态
	BoothId     uint64  `json:"boothId" gorm:"booth_id"`               // 摊位id
	Model
}

type waresModel struct{}

var WaresModel waresModel
