package models

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/hoomaac/vertex/common"
	"github.com/hoomaac/vertex/common/vtypes"
	"github.com/hoomaac/vertex/middleware/jwt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	UID      uint64 `json:"uid" gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"name"`
	UserName string `json:"user_name" gorm:"unique;"`
	RCode    string `gorm:"foreignKey:Referral" json:"r_code"`
}

func CreateUser(user *User) (int, *vtypes.AuthResponse) {

	db := common.Db

	var mySqlErr *mysql.MySQLError

	if len(user.UserName) == 0 || len(user.Name) == 0 {
		return vtypes.BadRequest, &vtypes.AuthResponse{Status: vtypes.BadRequest, Data: "username or name is empty"}
	}

	newUser := &User{
		Name:     user.Name,
		UserName: user.UserName,
		RCode:    user.RCode,
	}

	trx := db.Create(newUser)

	if errors.As(trx.Error, &mySqlErr) &&  mySqlErr.Number == common.DuplicateMysqlErr{
		return vtypes.BadRequest, &vtypes.AuthResponse{Status: vtypes.BadRequest, Data: "this username is already registered"}
	}

	token := jwt.GenerateJwt(newUser.UserName, newUser.Name)

	return vtypes.Ok, &vtypes.AuthResponse{Status: vtypes.Ok, Data: token}
}
