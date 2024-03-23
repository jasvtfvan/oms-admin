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
		r.GET("checkinit", dbApi.CheckInit) // 检查是否需要初始化
		r.GET("initdb", dbApi.InitDB)       // 初始化db
	}
}

func (*DbRouter) InitDbPrivateRouter(router *gin.RouterGroup) {
	r := router.Group("init")
	dbApi := v1.ApiGroupApp.System.DbApi
	{
		r.GET("checkupdate", dbApi.CheckUpdate) // 检查更新
		r.GET("updatedb", dbApi.UpdateDB)       // 更新db
	}
}
