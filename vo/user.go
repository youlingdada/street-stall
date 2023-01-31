package vo

import "time"

// UserVo 用户模型
type UserVo struct {
	UId       uint64    `json:"uId"`       // 用户id
	Username  string    `json:"username"`  // 用户名
	Email     string    `json:"email"`     //用户邮箱
	Phone     string    `json:"phone"`     // 用户电话
	Type      uint8     `json:"userType"`  // 用户类型 0 商户（卖家） 1 用户（买家）
	Status    uint8     `json:"status"`    // 用户状态
	AvatarUrl string    `json:"avatarUrl"` // 用户头像链接
	CreateAt  time.Time `json:"createAt"`  // 创建时间
	UpdateAt  time.Time `json:"updateAt"`  // 更新时间
}
