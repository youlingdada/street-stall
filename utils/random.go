/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-08 14:14:46
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-08 15:01:33
 * @FilePath: \street-stall\utils\random.go
 * @Description：生产相关的随机数据
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */

package utils

import (
	"math/rand"
	"time"
)

var randomData []byte = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z'}

/**
 * @description: 获取随机字符
 * @param {int} count 字符个数
 * @return {*}
 */
func GetRandomCode(count int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := ""
	for i := 0; i < count; i++ {
		res = res + string(randomData[r.Intn(len(randomData))])
	}
	return res
}

/**
 * @description: 获取随机字符，纯数字
 * @param {int} count 字符个数
 * @return {*}
 */
func GetRandomCodeOnlyNumber(count int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := ""
	for i := 0; i < count; i++ {
		res = res + string(randomData[r.Intn(10)])
	}
	return res
}

/**
 * @description: 获取随机的字符，纯字符
 * @param {int} count 字符个数
 * @return {*}
 */
func GetRandomCodeOnlyChar(count int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := ""
	for i := 0; i < count; i++ {
		res = res + string(randomData[r.Intn(len(randomData)-10)+10])
	}
	return res
}
