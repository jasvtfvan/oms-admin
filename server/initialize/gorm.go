package initialize

import (
	"os"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/demo"
	"github.com/jasvtfvan/oms-admin/server/model/system"
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
	// 自动更新的表结构切片
	var Tables = []interface{}{
		&system.SysUser{},
		&system.SysGroup{},
		&system.SysRole{},
		&demo.Demo{},
		// 添加其他需要迁移的表结构
	}

	db := global.OMS_DB
	err := db.AutoMigrate(
		Tables...,
	)
	if err != nil {
		global.OMS_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.OMS_LOG.Info("register table success")
}
