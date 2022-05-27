package models

type Goods struct {
	GID  uint64 `json:"gid" gorm:"column:GID;primary_key;auto_increment;not_null"`
	Name string `json:"name"`
}

type Good struct {
	Name string `json:"name" binding:"required"`
}
