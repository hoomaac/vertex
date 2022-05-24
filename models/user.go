package models

import "github.com/jinzhu/gorm"

var User struct {
	gorm.Model

	UID      uint64 `json:"uid" gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	RCode    string `gorm:"foreignKey:Referral" json:"r_code"`
}
