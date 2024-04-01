package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasvtfvan/oms-admin/server/api/v1"
)

type CacheRouter struct{}

func (*CacheRouter) InitCachePublicRouter(router *gin.RouterGroup) {}

func (*CacheRouter) InitCachePrivateRouter(router *gin.RouterGroup) {}

func (*CacheRouter) InitCacheCasbinRouter(router *gin.RouterGroup) {
	r := router.Group("cache")
	cacheApi := v1.ApiGroupApp.System.CacheApi
	{
		r.POST("test-cache", cacheApi.DoTestCache)
	}
}
