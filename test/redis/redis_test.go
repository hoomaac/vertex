package test

import (
	"log"
	"os"
	"testing"

	"github.com/hoomaac/vertex/pkg/config"
	"github.com/hoomaac/vertex/pkg/redis"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Printf("failed, %v\n", err)
		os.Exit(1)
	}

	redisConf := &config.RedisConfig{
		Ip:       os.Getenv("REDIS_IP"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
	}

	redis.InitRedis(redisConf)

	if redis.RedisClient == nil {
		log.Println("failed to init redis")
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)
}

func TestSetGetValue(t *testing.T) {

	if !redis.SetValue("key", "pass@value", 20) {
		t.Error("redis set value failed")
	}

	value := redis.GetValue("key")

	if value == "" {
		t.Error("redis get value failed")
	}

}
