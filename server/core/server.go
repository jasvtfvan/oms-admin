package core

import (
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize"
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
	// initialize.Redis()

	// 从db加载jwt数据
	// 	if global.OMS_DB != nil {
	// 		system.LoadAll()
	// 	}

	router := initialize.Routers()
	// 	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.OMS_CONFIG.System.Addr)
	var s interface{}
	if global.OMS_CONFIG.System.UseTls {
		s = initServerTLS(address, router)
	} else {
		s = initServer(address, router)
	}

	global.OMS_LOG.Info("Server run success on ", zap.String("address", address))

	fmt.Printf(`
		欢迎使用 oms-admin
		当前版本:v2.6.0
		默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	`, address)

	if global.OMS_CONFIG.System.UseTls { // 开启https验证
		certFile, keyFile := global.OMS_CONFIG.System.TlsCert, global.OMS_CONFIG.System.TlsKey
		global.OMS_LOG.Error(s.(serverTLS).ListenAndServeTLS(certFile, keyFile).Error())
	} else {
		global.OMS_LOG.Error(s.(server).ListenAndServe().Error())
	}
}
