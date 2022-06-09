package models

import "gorm.io/gorm"

type Referral struct {
	gorm.Model

	RID   uint64 `json:"rid" gorm:"primary_key;auto_increment;not_null"`
	Code  string `json:"code"`
	Count uint64 `json:"count"`
}
