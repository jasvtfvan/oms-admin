package main

import (
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/core"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
	initCache "github.com/jasvtfvan/oms-admin/server/initialize/cache"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

// @title                       Oms-Admin Swagger API接口文档
// @version                     V1.0.0
// @description                 使用gin的全栈开发基础平台
// @host												127.0.0.1:8888
// @basePath                    /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @securityDefinitions.apikey  ApiKeyDomain
// @in                          header
// @name												x-group
func main() {
	core.WritePIDToFile()              // vscode debug模式下，将pid写入txt文件，停止时可以根据pid进行kill操作
	global.OMS_VP = core.Viper()       // 加载配置文件
	initialize.BaseInit()              // 验证基础信息
	global.OMS_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.OMS_LOG) // 使用全局log

	/*
		连接gorm数据库
	*/
	global.OMS_DB = initialize.Gorm() // gorm连接数据库 [导入initialize包，register_init执行]
	if global.OMS_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.OMS_DB.DB()
		defer func() {
			fmt.Println(utils.GetStringWithTime("====== [Golang] main.go 关闭DB连接 ======"))
			db.Close()
		}()
	}

	/*
		创建freecache本机缓存
	*/
	global.OMS_FREECACHE = initCache.GetFreecacheClient()
	/*
		连接redis缓存
	*/
	if global.OMS_CONFIG.System.AuthCache == "redis" {
		redisConfig := global.OMS_CONFIG.Redis
		global.OMS_REDIS = initCache.GetRedisClient(redisConfig) // 初始化redis服务
		if global.OMS_REDIS != nil && global.OMS_REDIS.Client != nil {
			rdb := global.OMS_REDIS.Client
			defer func() {
				fmt.Println(utils.GetStringWithTime("====== [Golang] main.go 关闭REDIS连接 ======"))
				rdb.Close()
			}()
		}
	}

	core.RunWindowsServer()
}
