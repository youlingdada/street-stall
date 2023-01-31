// Package api /*
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/service"
	"github.com/youlingdada/street-stall/utils"
	"github.com/youlingdada/street-stall/vo"
	"log"
)

type userApi struct{}

var UserApi userApi

// UserLogin 用户登录接口
// @Summary 用户登录
// @Schemes
// @Description 用户登录
// @Tags 用户模块
// @Param user body query.UserLoginQuery true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/login [post]
func (ua userApi) UserLogin(c *gin.Context) {
	var user query.UserLoginQuery
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("用户登录解析用户信息失败,err->%s\n", err)
		utils.DealResult(c, utils.Error("用户信息错误"))
	}

	// 获取用户登录后的token
	token, err := service.UserService.UserLogin(user)

	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("登录成功", token))
	}
}

// UserRegister 用户注册接口
// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags 用户模块
// @Param user body query.UserAddQuery true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/register [post]
func (ua userApi) UserRegister(c *gin.Context) {
	var q query.UserAddQuery
	err := c.BindJSON(&q)
	if err != nil {
		log.Printf("JSON绑定失败: err->%s\n", err)
		utils.DealResult(c, utils.Error("用户信息错误"))
		return
	}
	// 校验参数
	err = query.UserAddValidation(q)
	if err != nil {
		log.Printf("用户注册失败：err->%s\n", err)
		utils.DealResult(c, utils.Error(err.Error()))
	}
	if err := service.UserService.UserRegister(q); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("注册成功"))
	}
}

// UserUpdateInfo 用户信息修改接口
// @Summary 用户信息修改
// @Schemes
// @Description 用户信息修改接口
// @Tags 用户模块
// @Param user body query.UserUpdateInfoQuery true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/updateInfo [post]
func (ua userApi) UserUpdateInfo(c *gin.Context) {
	var infoQuery query.UserUpdateInfoQuery
	err := c.BindJSON(&infoQuery)
	if err != nil {
		return
	}

	if err := service.UserService.UserUpdateInfo(infoQuery); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("信息修改成功"))
	}
}

// UserUpdateAvatar 更新用户头像
// @Summary 更新用户头像
// @Schemes
// @Description 用户更新头像图片
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id query int64 true "文件id"
// @Success 200 {object} utils.Result
// @Router /user/updateAvatar [post]
func (ua userApi) UserUpdateAvatar(c *gin.Context) {
	if avatarUrl, err := service.UserService.UserUpdateAvatar(c); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("信息修改成功", avatarUrl))
	}
}

// UserSendRegEmailCode 获取注册邮箱验证码
// @Summary 获取注册邮箱验证码
// @Schemes
// @Description 获取注册邮箱验证码
// @Tags 用户模块
// @Param email query string true "用户邮箱"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/sendRegCode [get]
func (ua userApi) UserSendRegEmailCode(c *gin.Context) {
	email := c.Query("email")

	if err := service.UserService.UserSendRegEmailCode(email); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("发送验证码成功"))
	}
}

// GetLoginUser 获取登录的用户信息
// @Summary 获取登录的用户信息
// @Schemes
// @Description 在用户登录之后，可以获取其相关信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/getLoginUser [get]
func (ua userApi) GetLoginUser(c *gin.Context) {
	value, ok := c.Get("loginUser")
	if ok {
		loginUser := value.(vo.UserVo)
		utils.DealResult(c, utils.Success("获取登录信息成功", loginUser))
	} else {
		utils.DealResult(c, utils.Error("获取登录信息失败，请重新登录"))
	}
}

// ModifyUserPwd 用户修改密码
// @Summary 用户修改密码
// @Schemes
// @Description 需要邮箱验证，用户修改密码
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/getLoginUser [get]
func ModifyUserPwd() {

}

// UserUpdatePwd 用户修改密码
// @Summary 用户修改密码
// @Schemes
// @Description 用户修改密码
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body query.UserUpdatePwdQuery true "用户信息"
// @Success 200 {object} utils.Result
// @Router /user/updatePwd [post]
func (ua userApi) UserUpdatePwd(c *gin.Context) {
	var pwdQuery query.UserUpdatePwdQuery
	err := c.BindJSON(&pwdQuery)
	if err != nil {
		log.Printf("修改密码解析数据失败,email->%s,err->%s\n", pwdQuery.Email, err)
		utils.DealResult(c, utils.Error("修改密码失败，数据格式错误"))
		return
	}
	err = service.UserService.UserUpdatePwd(pwdQuery)
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("密码修改成功"))
	}
}

// UserSendModifyPwdEmail 获取更改密码邮箱验证码
// @Summary 获取更改密码邮箱验证码
// @Schemes
// @Description 获取更改密码邮箱验证码
// @Tags 用户模块
// @Param email query string true "用户邮箱"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/sendModifyCode [get]
func (ua userApi) UserSendModifyPwdEmail(c *gin.Context) {
	email := c.Query("email")
	if err := service.UserService.UserSendModifyPwdEmailCode(email); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("发送验证码成功"))
	}
}

// UserModifyUsername 用户更改昵称
// @Summary 用户更改昵称
// @Schemes
// @Description 用户更改昵称
// @Tags 用户模块
// @Param username query string true "用户昵称"
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /user/modifyUsername [post]
func (ua userApi) UserModifyUsername(c *gin.Context) {
	username := c.Query("username")
	if err := service.UserService.UserModifyUsername(c, username); err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.SuccessNoData("更新用户昵称成功"))
	}
}
