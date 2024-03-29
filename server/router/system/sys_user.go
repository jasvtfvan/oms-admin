package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasvtfvan/oms-admin/server/api/v1"
)

type UserRouter struct{}

func (*UserRouter) InitUserPublicRouter(router *gin.RouterGroup) {
	r := router.Group("base")
	userApi := v1.ApiGroupApp.System.UserApi
	{
		r.POST("login", userApi.Login)
		r.POST("captcha", userApi.Captcha)
	}
}

func (*UserRouter) InitUserPrivateRouter(router *gin.RouterGroup) {}