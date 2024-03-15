package initialize

import (
	"os"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/goods"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.OMS_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.OMS_DB
	err := db.AutoMigrate(
		&system.SysUser{},
		&system.SysGroup{},
		&system.SysRole{},
		&goods.GoodsOrder{},
	)
	if err != nil {
		// global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	// global.GVA_LOG.Info("register table success")
}
