package main

import (
	"github.com/jasvtfvan/oms-admin/server/core"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.6.0
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
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
	core.RunWindowsServer()
}
