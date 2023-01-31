/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-07 09:15:58
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-21 11:58:39
 * @FilePath: \street-stall\utils\rsa.go
 * @Description:
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */

package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
)

var publicKey *rsa.PublicKey
var privateKey *rsa.PrivateKey

func init() {
	private, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Printf("rsa init failure,err: %s\n", err)
		log.Panic("Application stop")
	}
	privateKey = private
	publicKey = &private.PublicKey
}

// GetRSAPublicKey /**
func GetRSAPublicKey() (*string, error) {
	pkixPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	res := base64.StdEncoding.EncodeToString(pkixPublicKey)
	return &res, nil
}

// RSADecode /**
func RSADecode(cipher string) (*string, error) {
	bytes, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return nil, err
	}
	temp, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, bytes)
	if err != nil {
		return nil, err
	}
	res := string(temp)
	return &res, err
}

//ARwBni42k1xCRibSVk+sA5n1et9K59Pz94gClbWGeZj/1vikARXVHe09/Pbs9nmucnRCOxebDesC46WYyQnXEloTIvfRy9qnBVfSvpAqLS7pV/J8ZuZL/w/tdL28gIK+JrcwscFMwRT1tazSlp6EcaJB2blOXltVs0M9uCw0LBA=
//ARwBni42k1xCRibSVk sA5n1et9K59Pz94gClbWGeZj/1vikARXVHe09/Pbs9nmucnRCOxebDesC46WYyQnXEloTIvfRy9qnBVfSvpAqLS7pV/J8ZuZL/w/tdL28gIK JrcwscFMwRT1tazSlp6EcaJB2blOXltVs0M9uCw0LBA=

// RsaEncode /**
func RsaEncode(plain string) (*string, error) {
	temp, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plain))
	if err != nil {
		return nil, err
	}
	res := base64.StdEncoding.EncodeToString(temp)
	return &res, err
}
