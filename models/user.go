package models

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/hoomaac/vertex/pkg/app"
	"github.com/hoomaac/vertex/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UID      uint64 `json:"uid" gorm:"primary_key;auto_increment;not_null"`
	Email    string `json:"email" gorm:"unique;"`
	UserName string `json:"user_name" gorm:"unique;"`
	RCode    string `gorm:"foreignKey:Referral" json:"r_code"`
	Verified bool   `json:"verified"`
}

func CreateUser(registerReg *app.RegisterRequest) *User {

	db := database.Db

	var mySqlErr *mysql.MySQLError

	if len(registerReg.Username) == 0 || len(registerReg.Email) == 0 {
		return nil
	}

	newUser := &User{
		Email:    registerReg.Email,
		UserName: registerReg.Username,
	}

	trx := db.Create(newUser)

	if errors.As(trx.Error, &mySqlErr) && mySqlErr.Number == database.DuplicateMysqlErr {
		return nil
	}

	return newUser
}

func FindUserByUsername(username string) *User {

	db := database.Db

	var user User

	db.Find(&user, "username = ?", username)

	return &user
}

func FindUserByEmail(email string) *User {

	db := database.Db

	var user User

	db.Find(&user, "email = ?", email)

	return &user
}

func UpdateUser(user *User) bool {

	db := database.Db

	

	err := db.Model(user).Update("verified", true).Error

	if err != nil {
		return false
	}

	return true
}
