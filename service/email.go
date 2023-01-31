/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-07 17:18:50
 * @LastEditors: youlingdada cxyyxd1224721
 * @LastEditTime: 2022-07-24 16:22:47
 * @FilePath: \street-stall\service\email.go
 * @Description:
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */
package service

import (
	"errors"
	"github.com/youlingdada/street-stall/utils/redis_utils"
	"log"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/youlingdada/street-stall/config"
	"github.com/youlingdada/street-stall/utils"
)

/**
 * @description: 发送邮箱验证码
 * @param {utils.Email} email 邮件信息
 * @param {string} tag
 * @return {*}
 */
func SendEmailCode(email utils.Email, tag string) error {
	appConfig := config.GlobalConfig
	email.From = appConfig.Email.Host

	// 获取验证码
	code := utils.GetRandomCode(4)
	email.Context = code
	// 将所有的字母转为小写
	code = strings.ToLower(code)

	// 构造redis key
	key := tag + email.To

	exp := GetExpByTag(tag)

	if err := redis_utils.RedisUtils.Setex(key, code, exp); err != nil {
		log.Printf("将验证码储存到redis失败,code->%s,err->%s\n", code, err)
		return errors.New("储存验证码失败")
	}

	if err := utils.SendEmail(email); err != nil {
		log.Printf("发送邮箱验证码失败,email->%s,err->%s\n", email.To, err)
		return errors.New("发送验证码失败")
	}
	return nil
}

/**
 * @description: 验证邮验证码
 * @param {string} email 邮箱
 * @param {string} code 验证码
 * @param {string} tag 唯一标识,用于表示邮箱验证码的作用
 * @return {*} 是否验证成功
 */
func VerifyEmailCode(email string, code, tag string) error {

	// 验证码转为小写
	code = strings.ToLower(code)

	key := tag + email
	ttl, err := redis_utils.RedisUtils.Ttl(key)
	if err != nil {
		return err
	}
	if ttl <= 0 {
		return errors.New("验证码已经过期")
	}

	// 从redis 中获取生成的验证码
	enCode, err := redis.String(redis_utils.RedisUtils.Get(key))
	if err != nil {
		return err
	}

	// 返回比对结果
	if code != enCode {
		return errors.New("邮箱验证码不匹配")
	}
	return nil
}

func GetExpByTag(tag string) int {
	switch tag {
	case utils.REG_EMAIL: // 注册验证码
		return 30 * 60
	case utils.UPDATE_EMAIL_EMAIL: // 更新邮箱验证码
		return 30 * 60
	case utils.UPDATE_PHONE_EMAIL: // 更新电话验证码
		return 30 * 60
	case utils.UPDATE_PWD_EMAIL: // 更新密码验证码
		return 30 * 60
	default:
		return 30 * 60 // 默认 30s * 60 = 30m
	}
}
