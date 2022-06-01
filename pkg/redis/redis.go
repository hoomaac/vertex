package redis

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hoomaac/vertex/pkg/config"
)

var lock = &sync.Mutex{}

var ctx = context.Background()

var RedisClient *redis.Client

// InitRedis initialises a global instance of redis.Client.
// This is a thread safe function.
func InitRedis(RedisConf *config.RedisConfig) {

	if RedisClient != nil {
		return
	}

	lock.Lock()

	// unlock the lock after an instance of RedisClient is initialised
	defer lock.Unlock()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RedisConf.Ip, RedisConf.Port),
		Password: RedisConf.Password,
		DB:       RedisConf.Database,
	})
}

func SetValue(key string, value interface{}, timeout time.Duration) bool {

	err := RedisClient.Set(ctx, key, value, timeout).Err()

	if err != nil {
		log.Printf("set key=%s, value=%v to redis failed due to %v", key, value, err)
		return false
	}

	return true
}

func GetValue(key string) string {

	value, err := RedisClient.Get(ctx, key).Result()

	if err != nil {
		log.Printf("get key=%s, value=%v from redis failed due to %v", key, value, err)
		return ""
	}

	return value
}
