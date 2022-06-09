package models

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model

	VID  uint64 `json:"vid" gorm:"primary_key;auto_increment;not_null"`
	Code string `json:"code"`
	UID  uint64 `json:"uid" gorm:"foreignKey:User"`
	GID  uint64 `json:"gid" gorm:"foreignKey:Good"`
}
