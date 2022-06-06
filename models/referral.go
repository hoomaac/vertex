package models

import (
	"github.com/hoomaac/vertex/common"
	"github.com/jinzhu/gorm"
)

type Referral struct {
	gorm.Model

	RID   uint64 `json:"rid" gorm:"primary_key;auto_increment;not_null"`
	Code  string `json:"code"`
	Count uint64 `json:"count"`
}

func GetAllReferral(referals *[]Referral) (err error) {
	db := common.Db
	if err = db.Model(&Referral{}).Find(referals).Error; err != nil {
		return err
	}
	return nil
}

func GetReferral(referral *Referral, id string) (err error) {
	db := common.Db
	if err = db.Model(&Referral{}).Where("rid = ?", id).First(referral).Error; err != nil {
		return err
	}
	return nil
}

func AddReferral(referral *Referral) (err error) {
	db := common.Db
	if err = db.Model(&Referral{}).Create(referral).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReferral(referral *Referral, id string) (err error) {
	db := common.Db
	if err = db.Save(&referral).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReferral(referral *Referral, id string) (err error) {
	db := common.Db
	if err = db.Model(&Referral{}).Where("rid = ?", id).Delete(referral).Error; err != nil {
		return err
	}
	return nil
}
