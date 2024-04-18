package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasvtfvan/oms-admin/server/api/v1"
)

type DbRouter struct{}

func (*DbRouter) InitDbPublicRouter(router *gin.RouterGroup) {
	r := router.Group("init")
	dbApi := v1.ApiGroupApp.System.DbApi
	{
		r.POST("check", dbApi.CheckInit) // 检查是否需要初始化
		r.POST("db", dbApi.InitDB)       // 初始化db
	}
}

func (*DbRouter) InitDbPrivateRouter(router *gin.RouterGroup) {
	r := router.Group("update")
	dbApi := v1.ApiGroupApp.System.DbApi
	{
		r.POST("check", dbApi.CheckUpdate) // 检查更新
		r.POST("db", dbApi.UpdateDB)       // 更新db
	}
}

func (*DbRouter) InitDbCasbinRouter(router *gin.RouterGroup) {}
