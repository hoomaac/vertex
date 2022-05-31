package models

type Good struct {
	GID  uint64 `json:"gid" gorm:"column:GID;primary_key;auto_increment;not_null"`
	Name string `json:"name"`
}
