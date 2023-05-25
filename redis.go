package core

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisHelper(config *RedisConfig) *RedisClientHelper {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
	})
	return &RedisClientHelper{
		client: client,
		ctx:    context.Background(),
	}
}

type RedisClientHelper struct {
	client *redis.Client
	ctx    context.Context
}

// 释放资源
func (rh *RedisClientHelper) Release() {
	rh.client.Close()
}
func (rh *RedisClientHelper) timeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func (rh *RedisClientHelper) Ping() (string, error) {
	ctx, cancel := rh.timeoutContext(time.Second * 5)
	defer cancel()
	return rh.client.Ping(ctx).Result()

}
func (rh *RedisClientHelper) Keys(pattern string) (bool, []string) {
	result, err := rh.client.Keys(rh.ctx, pattern).Result()
	if nil != err {
		return false, nil
	}
	return true, result
}

// 带过期时间
func (rh *RedisClientHelper) SetEx(key string, value any, expiration time.Duration) bool {
	result, err := rh.client.Set(rh.ctx, key, value, expiration).Result()
	if nil != err {
		return false
	}
	return result == "OK"
}
func (rh *RedisClientHelper) Set(key string, value any) bool {
	return rh.SetEx(key, value, 0)
}

func (rh *RedisClientHelper) Delete(keys ...string) int64 {
	result, err := rh.client.Del(rh.ctx, keys...).Result()
	if nil != err {
		return -1
	}
	return result
}

func (rh *RedisClientHelper) Expire(key string, expiration time.Duration) bool {
	result, err := rh.client.Expire(rh.ctx, key, expiration).Result()
	if nil != err {
		return false
	}
	return result
}

func (rh *RedisClientHelper) GetString(key string) (bool, string) {
	result, err := rh.client.Get(rh.ctx, key).Result()
	if nil != err {
		return false, ""
	}
	return true, result

}
func (rh *RedisClientHelper) GetInt(key string) (bool, int) {
	result, err := rh.client.Get(rh.ctx, key).Int()
	if nil != err {
		return false, 0
	}
	return true, result

}
func (rh *RedisClientHelper) GetInt64(key string) (bool, int64) {
	result, err := rh.client.Get(rh.ctx, key).Int64()
	if nil != err {
		return false, 0
	}
	return true, result

}
func (rh *RedisClientHelper) GetUint64(key string) (bool, uint64) {
	result, err := rh.client.Get(rh.ctx, key).Uint64()
	if nil != err {
		return false, 0
	}
	return true, result
}
func (rh *RedisClientHelper) GetFloat32(key string) (bool, float32) {
	result, err := rh.client.Get(rh.ctx, key).Float32()
	if nil != err {
		return false, 0
	}
	return true, result
}

func (rh *RedisClientHelper) GetFloat64(key string) (bool, float64) {
	result, err := rh.client.Get(rh.ctx, key).Float64()
	if nil != err {
		return false, 0
	}
	return true, result
}

func (rh *RedisClientHelper) GetBytes(key string) (bool, []byte) {
	result, err := rh.client.Get(rh.ctx, key).Bytes()
	if nil != err {
		return false, nil
	}
	return true, result
}
func (rh *RedisClientHelper) GetTime(key string) (bool, time.Time) {
	result, err := rh.client.Get(rh.ctx, key).Time()
	if nil != err {
		return false, time.Now()
	}
	return true, result
}

//================列表

func (rh *RedisClientHelper) LPush(key string, values ...any) int64 {

	result, err := rh.client.LPush(rh.ctx, key, values...).Result()
	if nil != err {
		return -1
	}
	return result
}
func (rh *RedisClientHelper) LPop(key string) (bool, string) {
	result, err := rh.client.LPop(rh.ctx, key).Result()
	if nil != err {
		return false, ""
	}
	return true, result
}
func (rh *RedisClientHelper) RPush(key string, values ...any) int64 {
	result, err := rh.client.RPush(rh.ctx, key, values...).Result()
	if nil != err {
		return -1
	}
	return result
}
func (rh *RedisClientHelper) RPop(key string) (bool, string) {
	result, err := rh.client.RPop(rh.ctx, key).Result()
	if nil != err {
		return false, ""
	}
	return true, result
}
