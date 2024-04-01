package router

import "github.com/gin-gonic/gin"

// 注册路由-不做鉴权
func InitPublicRouter(publicGroup *gin.RouterGroup) {
	systemRouter := RouterGroupApp.System
	systemRouter.InitDbPublicRouter(publicGroup)
	systemRouter.InitUserPublicRouter(publicGroup)
	systemRouter.InitCachePublicRouter(publicGroup)
}

// 注册路由-做鉴权
func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	systemRouter := RouterGroupApp.System
	systemRouter.InitDbPrivateRouter(privateGroup)
	systemRouter.InitUserPrivateRouter(privateGroup)
	systemRouter.InitCachePrivateRouter(privateGroup)
}

// 注册路由-做鉴权-casbin接口权限
func InitCasbinRouter(casbinGroup *gin.RouterGroup) {
	systemRouter := RouterGroupApp.System
	systemRouter.InitDbCasbinRouter(casbinGroup)
	systemRouter.InitUserCasbinRouter(casbinGroup)
	systemRouter.InitCacheCasbinRouter(casbinGroup)
}
