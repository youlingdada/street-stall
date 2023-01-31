// Package service /*
package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/utils/redis_utils"
	"github.com/youlingdada/street-stall/vo"
	"log"
	"strconv"
	"time"

	"github.com/youlingdada/street-stall/model"
	"github.com/youlingdada/street-stall/utils"
)

type userService struct{}

var UserService userService

// UserLogin 用户登录
// query 登录参数
func (us userService) UserLogin(query query.UserLoginQuery) (*string, error) {
	var user model.User
	user.Username = query.Username
	user.Password = query.Password
	// 根据用户名查出用户信息
	loginUser, err := GetUserByUsername(user.Username)
	if err != nil {
		log.Printf("用户名不存在,username->%s,err->%s\n", user.Username, err)
		return nil, errors.New("账户或密码错误")
	}
	// rsa 解密密码，并与数据库中密码比对
	pwd, err := utils.RSADecode(user.Password)
	if err != nil {
		log.Printf("rsa 密文解析失败,password->%s,err->%s\n", user.Password, err)
		return nil, errors.New("密码解析失败")
	}
	// 匹配数据库密文
	err = utils.CryptMatch(*pwd, loginUser.Password)
	if err != nil {
		log.Printf("密码比对错误,password->%s,err->%s\n", *pwd, err)
		return nil, errors.New("账号或密码错误")
	}
	// 获取jwt token
	// 密码置为空
	var userVo vo.UserVo
	userVo.Status = loginUser.Status
	userVo.UId = loginUser.UId
	userVo.Username = loginUser.Username
	userVo.Type = loginUser.Type
	userVo.CreateAt = loginUser.CreateAt
	userVo.UpdateAt = loginUser.UpdateAt
	userVo.Email = loginUser.Email
	userVo.Phone = loginUser.Phone

	fileVo, err := FileService.Detail(loginUser.AvatarId)
	if err == nil {
		userVo.AvatarUrl = fileVo.FileUrl
	}
	token, err := us.doLogin(userVo)
	if err != nil {
		log.Printf("生成jwt token 失败,loginUser->%s,err->%s\n", loginUser.UserToString(), err)
		return nil, errors.New("登录失败,生成用户信息token失败")
	}
	return token, nil
}

// UserRegister 用户注册
func (us userService) UserRegister(query query.UserAddQuery) error {
	var user model.User
	err := VerifyEmailCode(query.Email, query.Code, utils.REG_EMAIL)
	// 验证用户邮箱
	if err != nil {
		log.Printf("用户邮箱验证失败，邮箱->%s, err->%s\n", query.Email, err)
		return errors.New("验证码错误")
	}
	user.Email = query.Email
	user.Username = query.Username

	// 对相关加密的字段解密
	pwd, err := utils.RSADecode(query.Password)
	if err != nil {
		log.Printf("对密码rsa 解密失败, pwd->%s,err->%s\n", query.Password, err)
		return errors.New("用户注册失败")
	}
	tempPwd, err := utils.CryptEncode(*pwd)
	if err != nil {
		log.Printf("密码加密失败,password->%s\n", query.Password)
		return errors.New("用户注册失败")
	}
	user.Password = *tempPwd
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	tx := model.DB.Create(&user)
	err = tx.Error
	if err != nil {
		log.Printf("数据库插入用户信息失败, user->%s,err->%s\n", user.UserToString(), err)
		return errors.New("用户注册失败")
	}
	return nil
}

// UserUpdateInfo 更新用户信息
func (us userService) UserUpdateInfo(infoQuery query.UserUpdateInfoQuery) error {
	var user model.User
	user.Username = infoQuery.Username
	user.AvatarId = infoQuery.AvatarId

	err := model.UserUpdate(user)
	if err != nil {
		log.Printf("用户信息更新失败,user->%s,err->%s\n", user.UserToString(), err)
		return errors.New("用户信息更新失败")
	}
	return nil
}

// UserUpdateAvatar 更新用户头像
func (us userService) UserUpdateAvatar(c *gin.Context) (string, error) {
	loginUser := utils.GetLoginUser(c)
	id, _ := strconv.Atoi(c.Query("id"))
	fileVo, err := FileService.Detail(uint64(id))
	if err != nil {
		log.Printf("头像信息获取失败，err->%s\n", err)
		return "", errors.New("头像信息获取失败")
	}
	err = model.UserUpdateAvatar(loginUser.UId, fileVo.Id)
	if err != nil {
		log.Printf("用户信息更新失败,uId->%d,err->%s\n", loginUser.UId, err)
		return "", errors.New("用户头像更新失败")
	}
	value, _ := c.Get("uuid")
	key := value.(string)
	user, err := redis_utils.UserCache.Get(key)
	if err != nil {
		log.Printf("用户获取登录信息失败,err->%s\n", err)
		return "", errors.New("获取当前用户登录信息失败，请重新登录")
	}
	user.AvatarUrl = fileVo.FileUrl
	ok, err := redis_utils.UserCache.PutSecondsTtl(key, *user, redis_utils.UserExp)
	if err != nil || !ok {
		log.Printf("更新用户登录信息失败，请重新登录,err->%s\n", err)
		return "", errors.New("更新用户登录信息失败，请重新登录")
	}
	return fileVo.FileUrl, nil
}

/**
 * @description: 根据用户id 获取用户信息
 * @param {int64} userId 用户id
 * @return {*} 用户信息
 */
func GetUserById(userId uint64) (*model.User, error) {
	user, err := model.UserSelectOneById(userId)
	if err != nil {
		log.Printf("根据id获取用户失败,userId->%d,err->%s", userId, err)
		return nil, errors.New("用户id不存在")
	}
	return user, nil
}

/**
 * @description: 根据用户名获取用户信息
 * @param {string} username 用户名
 * @return {*} 用户信息
 */
func GetUserByUsername(username string) (*model.User, error) {
	user, err := model.UserSelectOneByUsername(username)
	if err != nil {
		log.Printf("根据username 获取用户信息失败,username->%s,err->%s\n", username, err)
		return nil, errors.New("用户昵称不存在")
	}
	return user, nil
}

/**
 * @description: 根据emai获取用户信息
 * @param {string} email
 * @return {*}
 */
func GetUserByEmail(email string) (*model.User, error) {
	user, err := model.UserSelectOneByEmail(email)
	if err != nil {
		log.Printf("根据email 获取用户信息失败,err->%s\n", err)
		return nil, errors.New("获取用户信息失败")
	}
	return user, nil
}

/**
 * @description: 根据phone 获取用户信息
 * @param {string} phone
 * @return {*}
 */
func GetUserByPhone(phone string) (*model.User, error) {
	user, err := model.UserSelectOneByPhone(phone)
	if err != nil {
		log.Printf("根据phone 获取用户信息失败,err->%s\n", err)
		return nil, errors.New("获取用户信息失败")
	}
	return user, nil
}

/**
 * @description: 检查用户名的唯一性
 * @param {string} username
 * @return {*}
 */
func UserCheckOnlyUsername(username string) error {
	_, err := GetUserByUsername(username)
	if err != nil {
		return nil
	}
	return errors.New("用户名已存在")
}

/**
 * @description: 检查邮箱的唯一性
 * @param {string} email
 * @return {*}
 */
func UserCheckOnlyEmail(email string) error {
	_, err := GetUserByEmail(email)
	if err != nil {
		return nil
	}
	return errors.New("邮箱已被占用")
}

/**
 * @description: 检查电话的唯一性
 * @param {string} phone
 * @return {*}
 */
func UserCheckOnlyPhone(phone string) error {
	_, err := GetUserByPhone(phone)
	if err != nil {
		return nil
	}
	return errors.New("电话号码已被注册")
}

func UserUpdateEmail(user model.User, code string) error {

	// 验证用户邮箱
	if err := VerifyEmailCode(user.Email, code, utils.UPDATE_EMAIL_EMAIL); err != nil {
		log.Printf("用户邮箱验证失败，邮箱->%s,err->%s\n", user.Email, err)
		return errors.New("邮箱验证失败")
	}

	if err := model.UserUpdate(user); err != nil {
		log.Printf("修改邮箱失败,email->%s,err->%s", user.Email, err)
		return errors.New("邮箱修改失败")
	}
	return nil
}

func UserUpdatePhone(user model.User, code string) error {

	// 验证用户邮箱
	if err := VerifyEmailCode(user.Email, code, utils.UPDATE_PHONE_EMAIL); err != nil {
		log.Printf("用户邮箱验证失败，邮箱->%s, err->%s\n", user.Email, err)
		return errors.New("邮箱验证失败")
	}

	if err := model.UserUpdate(user); err != nil {
		log.Printf("修改电话失败,phone->%s,err->%s", user.Phone, err)
		return errors.New("电话号码修改失败")
	}
	return nil
}

// UserUpdatePwd 更新用户密码
func (us userService) UserUpdatePwd(query query.UserUpdatePwdQuery) error {
	// 验证用户邮箱
	if err := VerifyEmailCode(query.Email, query.Code, utils.UPDATE_PWD_EMAIL); err != nil {
		log.Printf("用户邮箱验证失败，邮箱->%s, err->%s\n", query.Email, err)
		return errors.New("邮箱验证失败")
	}
	// 解密rsa
	pwd, err := utils.RSADecode(query.Password)
	if err != nil {
		log.Printf("rsa解密失败，email->%s,err->%s\n", query.Email, err)
		return errors.New("密码格式错误")
	}
	enPwd, err := utils.CryptEncode(*pwd)
	if err != nil {
		log.Printf("加密密码失败，email->%s,err->%s\n", query.Email, err)
		return errors.New("密码存储错误")
	}
	if err := model.UserUpdatePwd(query.Email, *enPwd); err != nil {
		log.Printf("修改密码失败,email->%s,err->%s", query.Email, err)
		return errors.New("密码修改失败")
	}
	return nil
}

// UserSendRegEmailCode 发送注册用户时验证码
func (us userService) UserSendRegEmailCode(email string) error {
	var e utils.Email
	e.To = email
	e.Subject = "用户注册验证码"

	if err := SendEmailCode(e, utils.REG_EMAIL); err != nil {
		log.Printf("发送验证码失败,email->%s,err->%s\n", email, err)
		return errors.New("发送验证码失败")
	}
	return nil
}

// UserSendModifyPwdEmailCode 发送修改密码时的邮箱验证码
func (us userService) UserSendModifyPwdEmailCode(email string) error {
	var e utils.Email
	e.To = email
	e.Subject = "用户更改密码"
	if err := SendEmailCode(e, utils.UPDATE_PWD_EMAIL); err != nil {
		log.Printf("发送验证码失败,email->%s,err->%s\n", email, err)
		return errors.New("发送验证码失败")
	}
	return nil
}

// UserModifyUsername 更新昵称
func (us userService) UserModifyUsername(c *gin.Context, username string) error {
	loginUser := utils.GetLoginUser(c)

	err := model.UserUpdateUsername(loginUser.UId, username)
	if err != nil {
		log.Printf("用户更新昵称失败,uId->%d,err->%s", loginUser.UId, err)
		return errors.New("用户更改昵称失败")
	}
	return nil
}

// 登录流程，返回token
func (us userService) doLogin(user vo.UserVo) (*string, error) {
	userCache := redis_utils.UserCache
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	ok, err := userCache.PutSecondsTtl(newUUID.String(), user, redis_utils.UserExp)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("登录用户缓存失败")
	}
	// 获取到token
	token, err := utils.GetJwtToken(newUUID.String())
	if err != nil {
		return nil, err
	}
	return token, nil
}
