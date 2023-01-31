package query

import "errors"

// UserAddQuery 用户注册模型
// @Description 用户注册数据
type UserAddQuery struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 用户密码
	Email    string `json:"email"`    //用户邮箱
	Phone    string `json:"phone"`    // 用户电话
	Type     uint8  `json:"userType"` // 用户类型 0 商户（卖家） 1 用户（买家）默认0
	Code     string `json:"code"`     // 邮箱验证码
}

// UserLoginQuery 用户登录模型
// @Description 用户登录数据
type UserLoginQuery struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 用户密码
	LoginType int8   `json:"loginType"` // 登录方式 0 用户名(默认) 1 邮箱 2 电话
}

// ModifyUserPwdQuery 用户修改密码模型
// @Description 用户修改密码模型
type ModifyUserPwdQuery struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 用户新密码
	Code     string `json:"code"`     // 邮箱验证码
}

// UserUpdateInfoQuery 用户信息更新模型
// @Description 用户信息更新模型
type UserUpdateInfoQuery struct {
	Username string `json:"username"` // 用户名
	AvatarId uint64 `json:"avatarId"` // 用户头像链接
}

// UserUpdatePwdQuery 用户密码更新模型
// @Description 用户修改密码模型
type UserUpdatePwdQuery struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 用户密码
	Code     string `json:"code"`     // 邮箱验证码
}

// UserAddValidation 验证用户添加参数
func UserAddValidation(query UserAddQuery) error {
	if query.Password == "" {
		return errors.New("用户密码不能为空")
	}
	if query.Username == "" {
		return errors.New("用户昵称不能为空")
	}
	if query.Email == "" {
		return errors.New("用户邮箱不能为空")
	}
	if query.Code == "" {
		return errors.New("用户验证码不能为空")
	}
	return nil
}
