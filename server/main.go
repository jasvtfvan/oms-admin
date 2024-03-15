package main

import (
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/core"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
	"go.uber.org/zap"
)

func main() {
	global.OMS_VP = core.Viper()           // 加载配置文件
	global.OMS_LOG = core.Zap()            // 初始化zap日志库
	zap.ReplaceGlobals(global.OMS_LOG)     // 使用全局log
	global.OMS_DB = initialize.GormMysql() // gorm连接数据库
	if global.OMS_DB != nil {
		initialize.RegisterTables() // 初始化表结构
		// 程序结束前关闭数据库链接
		db, _ := global.OMS_DB.DB()
		defer db.Close()
	}
	fmt.Println("hello admin")
}
