package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/api"
	"github.com/hoomaac/vertex/models"
	"github.com/hoomaac/vertex/pkg/config"
	"github.com/hoomaac/vertex/pkg/database"
	"github.com/hoomaac/vertex/pkg/redis"
	"github.com/joho/godotenv"
)

func setupDb() {

	dbInfo := &config.DataBaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Ip:       os.Getenv("DB_IP"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Fatalf("redisdb is mistyped, %v\n", err)
	}

	redisConf := &config.RedisConfig{
		Ip:       os.Getenv("REDIS_IP"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: redisDb,
	}

	database.InitDb(dbInfo)

	redis.InitRedis(redisConf)

	// migrate each model here
	database.Migrate(&models.User{})
	database.Migrate(&models.Good{})
}

func main() {

	// Start the engine and read configs from .env file
	engine := api.StartEngine(gin.DebugMode)

	env := godotenv.Load()

	if env != nil {
		log.Fatal("could not read env file")
	}

	port := os.Getenv("VERTEX_PORT")
	addr := os.Getenv("VERTEX_ADDR")

	setupDb()

	engine.Run(fmt.Sprintf("%s:%s", addr, port))
}
