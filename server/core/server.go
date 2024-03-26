package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
	"github.com/jasvtfvan/oms-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

type serverTLS interface {
	ListenAndServeTLS(certFile, keyFile string) error
}

func RunWindowsServer() {
	// 初始化redis服务
	initialize.Redis()

	// 从db加载jwt数据
	if global.OMS_DB != nil {
		system.LoadAll()
	}

	var router *gin.Engine
	// 开启zap的debug，才打印请求信息
	if global.OMS_CONFIG.Zap.Level == "debug" {
		logger := ZapGin()
		router = initialize.Routers(logger)
	} else {
		router = initialize.Routers(nil)
	}
	// 	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.OMS_CONFIG.System.Addr)
	var s interface{}
	if global.OMS_CONFIG.System.UseTls {
		s = initServerTLS(address, router)
	} else {
		s = initServer(address, router)
	}

	global.OMS_LOG.Info("Server run success on ", zap.String("address", address))

	version := global.OMS_CONFIG.Version
	fmt.Printf(`
		欢迎使用 oms-admin
		当前版本: %s
		默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	`, version, address)

	if global.OMS_CONFIG.System.UseTls { // 开启https验证
		certFile, keyFile := global.OMS_CONFIG.System.TlsCert, global.OMS_CONFIG.System.TlsKey
		global.OMS_LOG.Error(s.(serverTLS).ListenAndServeTLS(certFile, keyFile).Error())
	} else {
		global.OMS_LOG.Error(s.(server).ListenAndServe().Error())
	}
}
