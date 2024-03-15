package main

import (
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/core"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
)

func main() {
	global.OMS_VP = core.Viper()           // 加载配置文件
	initialize.CreateDatabase()            // 创建数据库
	global.OMS_DB = initialize.GormMysql() // gorm连接数据库
	if global.OMS_DB != nil {
		initialize.RegisterTables() // 初始化表结构
		// 程序结束前关闭数据库链接
		db, _ := global.OMS_DB.DB()
		defer db.Close()
	}
	fmt.Println("hello admin")
}
