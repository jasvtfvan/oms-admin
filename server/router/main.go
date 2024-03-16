package router

import "github.com/gin-gonic/gin"

// 注册路由-不做鉴权
func InitPublicRouter(publicGroup *gin.RouterGroup) {
	systemRouter := RouterGroupApp.System
	systemRouter.InitDbPublicRouter(publicGroup)
}

// 注册路由-做鉴权
func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	systemRouter := RouterGroupApp.System
	systemRouter.InitDbPrivateRouter(privateGroup)
}
