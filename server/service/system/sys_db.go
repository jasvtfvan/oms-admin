package system

import (
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"gorm.io/gorm"
)

type DBService interface {
	CheckDB() error
	CheckUpdate() error
}

type DBServiceImpl struct{}

// 检查数据连接
func (*DBServiceImpl) CheckDB() (err error) {
	if global.OMS_DB == nil {
		return errors.New("DB为空，数据库未创建")
	} else {
		db := global.OMS_DB
		tableCreated := db.Migrator().HasTable(&system.SysVersion{})
		if !tableCreated {
			return errors.New("表结构尚未创建")
		}
		err := db.Where("version_name = ?", "oms_version").First(&system.SysVersion{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("表数据尚未插入")
		}
	}
	return err
}

// 检查是否需要升级
func (*DBServiceImpl) CheckUpdate() (err error) {
	db := global.OMS_DB
	sysVersion := &system.SysVersion{}
	db.Where("version_name = ?", "oms_version").First(sysVersion)
	version := global.OMS_CONFIG.Version
	if sysVersion.Version != version {
		return errors.New("需升级为:" + version)
	}
	return err
}
