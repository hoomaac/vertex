package models

import (
	"github.com/jinzhu/gorm"
)

type Good struct {
	gorm.Model

	GID  uint64 `json:"gid" gorm:"primary_key;auto_increment;not_null"`
	Name string `json:"name"`
}
