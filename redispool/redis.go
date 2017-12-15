package redispool

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisPool struct {
	p *redis.Pool
}

func NewRedispool(url string) *RedisPool {
	redispool := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   800,
		IdleTimeout: 60 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url, redis.DialConnectTimeout(30*time.Second))
		},
	}
	return &RedisPool{redispool}
}
