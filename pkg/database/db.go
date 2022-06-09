package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/hoomaac/vertex/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define mysql error numbers
const (
	DuplicateMysqlErr = 1062
)

var Db *gorm.DB

func InitDb(database *config.DataBaseConfig) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username, database.Password, database.Ip, database.Port, database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("cannot open db, %v", err)
	}

	Db = db
}

// Create table each model
func Migrate(model interface{}) error {

	if Db == nil {
		return errors.New("Database is not initialized")
	}

	return Db.AutoMigrate(model)
}
