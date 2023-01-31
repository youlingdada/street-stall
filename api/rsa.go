/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-21 10:32:15
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-22 17:19:07
 * @FilePath: \street-stall\api\rsa.go
 * @Description:
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/utils"
	"log"
	"strings"
)

// GetPublicRsaKey 获取RSA加密公钥
// @Summary 获取加密公钥
// @Schemes
// @Description 获取RSA加密公钥
// @Tags 通用模块
// @Accept json
// @Produce json
// @Success 200 {object} utils.Result
// @Router /utils/getPublicRsaKey [get]
func GetPublicRsaKey(c *gin.Context) {
	publicKey, err := utils.GetRSAPublicKey()
	if err != nil {
		utils.DealResult(c, utils.Error(err.Error()))
	} else {
		utils.DealResult(c, utils.Success("获取rsa公钥成功", publicKey))
	}
}

// EncodeRsa 公钥加密数据
// @Summary 获取公钥的数据
// @Schemes
// @Description 获取RSA公钥加密的数据
// @Tags 通用模块
// @Accept json
// @Produce json
// @Param plain query string true "需要加密的数据"
// @Success 200 {object} utils.Result
// @Router /utils/encodeRsa [get]
func EncodeRsa(c *gin.Context) {
	plain := c.Query("plain")
	log.Printf("rsa 加密,plain->%s\n", plain)
	res, err := utils.RsaEncode(plain)

	if err != nil {
		utils.DealResult(c, utils.Error("rsa加密异常"))
	} else {
		utils.DealResult(c, utils.Success("rsa加密成功", res))
	}
}

// DecodeRsa 私钥解密数据
// @Summary 获取私钥解密数据
// @Schemes
// @Description 获取RSA私钥解密数据的数据
// @Tags 通用模块
// @Accept json
// @Produce json
// @Param cipher query string true "需要解密的数据"
// @Success 200 {object} utils.Result
// @Router /utils/decodeRsa [get]
func DecodeRsa(c *gin.Context) {
	cipher := c.Query("cipher")
	cipher = strings.Replace(cipher, " ", "+", len(cipher))
	res, err := utils.RSADecode(cipher)
	if err != nil {
		utils.DealResult(c, utils.Error("rsa解密异常"))
	} else {
		utils.DealResult(c, utils.Success("rsa解密成功", res))
	}
}
