package infra

import (
	"ARPostcard_server/biz/conf"
	"ARPostcard_server/biz/consts"
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
	"github.com/go-redis/redis"
	"time"
)

var RedisCli *redis.Client

func InitRedis() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.IP + ":" + string(rune(conf.Conf.Redis.Port)),
		Password: conf.Conf.Redis.Password,
		DB:       conf.Conf.Redis.Db,
		PoolSize: conf.Conf.Redis.PoolSize,
	})
}

// Set 添加键值对
func Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := RedisCli.Set(key, value, expiration).Err()
	if err != nil {
		ilog.EventError(ctx, err, "redis_set_err", "key", key, "value", value, "expiration", expiration)
		return ierror.NewIError(consts.DBError, "redis设置出错")
	}
	return err
}

// Match 判断键值对是否匹配
func Match(ctx context.Context, key string, value string) bool {
	val, err := RedisCli.Get(key).Result()
	if err != nil {
		ilog.EventError(ctx, err, "redis_set_err", "key", key, "value", value)
		return false
	}
	if val == value {
		return true
	}
	return false
}

func HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	val, err := RedisCli.HMSet(key, fields).Result()
	ilog.EventInfo(ctx, "redis_hmset", "key", key, "fields", fields, "val", val)
	if err != nil {
		ilog.EventError(ctx, err, "redis_set_err", "key", key, "value", val)
		return err
	}
	return nil
}

func HDel(ctx context.Context, key string, fields ...string) error {
	val, err := RedisCli.HDel(key, fields...).Result()
	ilog.EventInfo(ctx, "redis_hdel", "key", key, "fields", fields, "val", val)
	if err != nil {
		ilog.EventError(ctx, err, "redis_del_err", "key", key, "value", val)
		return err
	}
	return nil
}
