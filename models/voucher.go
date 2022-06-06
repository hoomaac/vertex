package models

import (
	"github.com/hoomaac/vertex/common"
	"github.com/jinzhu/gorm"
)

type Voucher struct {
	gorm.Model

	VID  uint64 `json:"vid" gorm:"primary_key;auto_increment;not_null"`
	Code string `json:"code"`
	UID  uint64 `json:"uid" gorm:"foreignKey:User"`
	GID  uint64 `json:"gid" gorm:"foreignKey:Good"`
}

func GetAllVouchers(vouchers *[]Voucher) (err error) {
	db := common.Db
	if err = db.Model(&Voucher{}).Find(vouchers).Error; err != nil {
		return err
	}
	return nil
}

func GetVoucher(voucher *Voucher, id string) (err error) {
	db := common.Db
	if err = db.Model(&Voucher{}).Where("vid = ?", id).First(voucher).Error; err != nil {
		return err
	}
	return nil
}

func AddVoucher(voucher *Voucher) (err error) {
	db := common.Db
	if err = db.Model(&Voucher{}).Create(voucher).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVoucher(voucher *Voucher, id string) (err error) {
	db := common.Db
	if err = db.Save(&voucher).Error; err != nil {
		return err
	}
	return nil
}

func DeleteVoucher(voucher *Voucher, id string) (err error) {
	db := common.Db
	if err = db.Model(&Voucher{}).Where("vid = ?", id).Delete(voucher).Error; err != nil {
		return err
	}
	return nil
}
