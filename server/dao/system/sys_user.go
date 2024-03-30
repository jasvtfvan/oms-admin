package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type UserDao struct{}

func UpdatePassword(id uint, password string) (int64, error) {
	db := global.OMS_DB
	res := db.Model(&system.SysUser{
		BaseModel: common.BaseModel{
			ID: id,
		},
	}).Update("password", password)
	return res.RowsAffected, res.Error
}

func EnableUser(id uint) (int64, error) {
	db := global.OMS_DB
	res := db.Model(&system.SysUser{
		BaseModel: common.BaseModel{
			ID: id,
		},
	}).Update("enable", true)
	return res.RowsAffected, res.Error
}

func DisableUser(id uint) (int64, error) {
	db := global.OMS_DB
	res := db.Model(&system.SysUser{
		BaseModel: common.BaseModel{
			ID: id,
		},
	}).Update("enable", false)
	return res.RowsAffected, res.Error
}

func DeleteUser(id uint) (int64, error) {
	db := global.OMS_DB
	res := db.Delete(&system.SysUser{}, id)
	return res.RowsAffected, res.Error
}

func FindUserById(id uint) (*system.SysUser, error) {
	var user system.SysUser
	db := global.OMS_DB
	err := db.First(&user, 10).Error
	return &user, err
}

func FindByUsername(username string) (*system.SysUser, error) {
	var user system.SysUser
	db := global.OMS_DB
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}
