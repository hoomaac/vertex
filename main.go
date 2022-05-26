package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoomaac/vertex/api"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/common/vtypes"
	"github.com/hoomaac/vertex/models"
	"github.com/joho/godotenv"
)

func setupDb() {

	dbInfo := &vtypes.DataBaseInfo{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Ip:       os.Getenv("DB_IP"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	common.InitDb(dbInfo)

	// migrate each model here
	common.Migrate(&models.User{})
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
