// Package model /*
package model

import (
	"encoding/json"
)

type User struct {
	UId      uint64 `json:"uId" gorm:"column:u_id;primaryKey"` // 用户id
	Username string `json:"username" gorm:"column:username"`   // 用户名
	Password string `json:"password" gorm:"column:password"`   // 用户密码
	Email    string `json:"email" gorm:"column:email"`         //用户邮箱
	Phone    string `json:"phone" gorm:"column:phone"`         // 用户电话
	Type     uint8  `json:"userType" gorm:"column:type"`       // 用户类型 0 商户（卖家） 1 用户（买家）
	Status   uint8  `json:"status" gorm:"column:status"`       // 用户状态
	AvatarId uint64 `json:"avatarId" gorm:"column:avatar_id"`  // 用户头像链接
	Model
}

// UserToString 把user对象转为字符串
func (u User) UserToString() string {
	res, _ := json.Marshal(u)
	return string(res)
}

// UserAdd 添加用户
func UserAdd(user User) error {
	res := DB.Create(&user)
	return res.Error
}

// UserDeleted 根据用户id删除用户
func UserDeleted(uId uint64) error {
	res := DB.Delete(&User{}, uId)
	return res.Error
}

// UserUpdate 更新用户信息
func UserUpdate(user User) error {
	res := DB.Save(user)
	return res.Error
}

// UserUpdateAvatar 更新用户头像
func UserUpdateAvatar(uId, fileId uint64) error {
	res := DB.Model(User{}).Where("u_id", uId).Update("avatar_id", fileId)
	return res.Error
}

// UserUpdatePwd 更新用户密码
func UserUpdatePwd(email, password string) error {
	res := DB.Model(User{}).Where("email", email).Update("password", password)
	return res.Error
}

// UserUpdateUsername 更新用户昵称
func UserUpdateUsername(uId uint64, username string) error {
	res := DB.Model(User{}).Where("u_id", uId).Update("username", username)
	return res.Error
}

// UserSelectOneById 根据用户id 查询用户
func UserSelectOneById(uId uint64) (*User, error) {
	var user User
	res := DB.First(&user, uId)
	return &user, res.Error
}

// UserSelectOneByUsername 根据用户昵称查询用户
func UserSelectOneByUsername(username string) (*User, error) {
	var user User
	res := DB.Where("username = ?", username).First(&user)
	return &user, res.Error
}

// UserSelectOneByEmail 根据用户邮箱查询用户
func UserSelectOneByEmail(email string) (*User, error) {
	var user User
	res := DB.Where("email = ?", email).First(&user)
	return &user, res.Error
}

// UserSelectOneByPhone 根据用户电话查询用户
func UserSelectOneByPhone(phone string) (*User, error) {
	var user User
	res := DB.Where("phone = ?", phone).First(&user)
	return &user, res.Error
}
