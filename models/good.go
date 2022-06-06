package models

import "github.com/hoomaac/vertex/common"

type Good struct {
	GID  uint64 `json:"gid" gorm:"column:GID;primary_key;auto_increment;not_null"`
	Name string `json:"name"`
}

func GetAllGoods(goods *[]Good) (err error) {
	db := common.Db
	if err = db.Model(&Good{}).Find(goods).Error; err != nil {
		return err
	}
	return nil
}

func GetGood(good *Good, id string) (err error) {
	db := common.Db
	if err = db.Model(&Good{}).Where("gid = ?", id).First(good).Error; err != nil {
		return err
	}
	return nil
}

func AddGood(good *Good) (err error) {
	db := common.Db
	if err = db.Model(&Good{}).Create(good).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGood(good *Good, id string) (err error) {
	db := common.Db
	if err = db.Save(&good).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGood(good *Good, id string) (err error) {
	db := common.Db
	if err = db.Model(&Good{}).Where("gid = ?", id).Delete(good).Error; err != nil {
		return err
	}
	return nil
}
