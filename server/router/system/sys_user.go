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

func (*UserRouter) InitUserPrivateRouter(router *gin.RouterGroup) {
	rb := router.Group("base")
	userApi := v1.ApiGroupApp.System.UserApi
	{
		rb.POST("compare-secret", userApi.CompareSecret)
	}
	ru := router.Group("user")
	{
		ru.GET("profile", userApi.GetUserProfile)
		ru.PUT("change-pwd", userApi.ChangePwd)
	}
}

func (*UserRouter) InitUserCasbinRouter(router *gin.RouterGroup) {
	r := router.Group("user")
	userApi := v1.ApiGroupApp.System.UserApi
	{
		r.DELETE("delete/:id", userApi.DeleteUser)
		r.PUT("disable/:id", userApi.DisableUser)
		r.PUT("enable/:id", userApi.EnableUser)
		r.PUT("reset-pwd", userApi.ResetPassword)
	}
}
