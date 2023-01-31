// Package model /*
package model

type Notice struct {
	NId       uint64 `json:"nId" gorm:"column:n_id;primaryKey"`  // 公告id
	Title     string `json:"title" gorm:"column:title"`          // 公告标题
	Content   string `json:"content" gorm:"column:content"`      // 公告内容
	AdminId   uint32 `json:"adminId" gorm:"column:admin_id"`     // 管理员id
	AdminName string `json:"adminName" gorm:"column:admin_name"` // 管理员昵称
	Model
}
