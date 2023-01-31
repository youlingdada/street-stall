package redis_utils

import (
	"encoding/json"
	"github.com/youlingdada/street-stall/vo"
)

type userRedisCache struct{}

var UserCache userRedisCache

const UserExp = 2 * 60 * 60

func (uc userRedisCache) Get(key string) (*vo.UserVo, error) {
	redisKey := uc.getPrefix() + key
	u, err := Conn.Do("Get", redisKey)
	if err != nil {
		return nil, err
	}
	user, err := uc.DescSerialize(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Put 加入缓存
func (uc userRedisCache) Put(key string, user vo.UserVo) (bool, error) {
	redisKey := uc.getPrefix() + key
	userJson, err := uc.Serialize(user)
	if err != nil {
		return false, err
	}
	reply, err := Conn.Do("set", redisKey, userJson)
	if err != nil {
		return false, err
	}
	// 判断是否设置成功
	status := reply.(string)
	if status == ReplyOk {
		return true, nil
	}
	return false, nil
}

// PutSecondsTtl 加入缓存
func (uc userRedisCache) PutSecondsTtl(key string, user vo.UserVo, ttl uint64) (bool, error) {
	redisKey := uc.getPrefix() + key
	userJson, err := uc.Serialize(user)
	if err != nil {
		return false, err
	}
	reply, err := Conn.Do("setex", redisKey, ttl, *userJson)
	if err != nil {
		return false, err
	}
	// 判断是否设置成功
	status := reply.(string)
	if status == ReplyOk {
		return true, nil
	}
	return false, nil
}

// Ttl 获取剩余时间
func (uc userRedisCache) Ttl(key string) (int64, error) {
	redisKey := uc.getPrefix() + key
	reply, err := Conn.Do("ttl", redisKey)
	if err != nil {
		return 0, err
	}
	return reply.(int64), nil
}

func (uc userRedisCache) getPrefix() string {
	return "USER_"
}

// Serialize 序列化user
func (uc userRedisCache) Serialize(user vo.UserVo) (*string, error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	ret := string(userJson)
	return &ret, nil
}

// DescSerialize 反序列化
func (uc userRedisCache) DescSerialize(userJson interface{}) (*vo.UserVo, error) {
	user := vo.UserVo{}
	data := string(userJson.([]uint8))
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
