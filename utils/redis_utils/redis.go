/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-08 11:03:14
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-12 09:29:26
 * @FilePath: \street-stall\utils\redis.go
 * @Description: redis 工具类
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */

package redis_utils

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/youlingdada/street-stall/config"
)

var Conn redis.Conn

const (
	ReplyOk = "OK"
)

func init() {
	app := config.GlobalConfig
	address := fmt.Sprintf("%s:%d", app.Redis.Host, app.Redis.Port)
	conn, err := redis.Dial("tcp", address, redis.DialDatabase(int(app.Redis.Db)), redis.DialUsername(app.Redis.Username), redis.DialPassword(app.Redis.Password))
	if err != nil {
		log.Printf("redis conn failure,err: %s", err)
		log.Panic("Application stop")
	}
	Conn = conn
}

type redisUtil struct{}

var RedisUtils redisUtil

/**
 * @description: redis设置值
 * @param {string} key 键
 * @param {interface{}} value 值
 * @return {*}
 */
func (ru redisUtil) Set(key string, value interface{}) error {
	_, err := Conn.Do("Set", key, value)
	return err
}

/**
 * @description: 设置值，并设置过期时间
 * @param {string} key 键
 * @param {interface{}} value 值
 * @param {int} expTime 过期时间 (s)
 * @return {*}
 */
func (ru redisUtil) Setex(key string, value interface{}, exp int) error {
	_, err := Conn.Do("Setex", key, exp, value)
	return err
}

/**
 * @description: redis 获取值
 * @param {string} key 键
 * @return {*} 返回的值
 */
func (ru redisUtil) Get(key string) (interface{}, error) {
	return Conn.Do("Get", key)
}

/**
 * @description: 获取指定键的过期时间
 * @param {string} key 键
 * @return {*} 过期时间
 */
func (ru redisUtil) Ttl(key string) (int64, error) {
	return redis.Int64(Conn.Do("Ttl", key))
}

// 删除一个指定的键值对
func (ru redisUtil) Del(key string) error {
	return nil
}
