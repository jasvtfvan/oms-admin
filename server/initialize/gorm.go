package initialize

import (
	"os"

	"github.com/jasvtfvan/oms-admin/server/global"
	"go.uber.org/zap"
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

// 如果程序修改了，表结构也跟着更新
func RegisterTables() {
	db := global.OMS_DB
	err := db.AutoMigrate(
		global.Tables...,
	)
	if err != nil {
		global.OMS_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.OMS_LOG.Info("register table success")
}
