package initialize

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	config := global.OMS_CONFIG
	switch config.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
