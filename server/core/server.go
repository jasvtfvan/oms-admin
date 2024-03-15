package core

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	// initialize.Redis()

	// 从db加载jwt数据
	// 	if global.OMS_DB != nil {
	// 		system.LoadAll()
	// 	}

	// 	Router := initialize.Routers()
	// 	Router.Static("/form-generator", "./resource/page")

	// 	address := fmt.Sprintf(":%d", global.OMS_CONFIG.System.Addr)
	// 	s := initServer(address, Router)

	// 	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	//	fmt.Printf(`
	//	欢迎使用 oms-admin
	//	当前版本:v2.6.0
	//	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	//
	// `, address)
	//
	//	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
