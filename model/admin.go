// Package model /*
package model

type Admin struct {
	AId       uint32 `json:"aId" gorm:"column:a_id;primaryKey"`  // 管理员id
	Username  string `json:"username" gorm:"column:username"`    // 管理员用户名
	Password  string `json:"password" gorm:"column:password"`    // 管理员密码
	AvatarUrl string `json:"avatarUrl" gorm:"column:avatar_url"` // 管理员头像链接
	Model
}
