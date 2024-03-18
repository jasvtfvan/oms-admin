package initialize

import (
	"fmt"

	_ "github.com/jasvtfvan/oms-admin/server/initialize/initializer/system"
)

// 只要main导入initialize包，init方法就执行
func init() {
	// do nothing,only import source package so that inits can be registered
	fmt.Println("====== [Golang] Init() 方法初始化 ======")
}
