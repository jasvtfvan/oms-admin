package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/service"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rootUsername := global.OMS_CONFIG.System.Username
		v := ctx.Value("claims")
		if claims, ok := v.(*utils.CustomClaims); ok {
			if claims.Username == rootUsername { // 超级管理员，所有接口都可以访问
				ctx.Next()
			} else {
				handler(ctx, claims)
			}
		} else {
			response.Fail(nil, "解析令牌信息失败", ctx)
			ctx.Abort()
			return
		}
	}
}

func handler(ctx *gin.Context, claims *utils.CustomClaims) {
	// 角色
	sub := ""
	// 域
	dom := ctx.Request.Header.Get("x-group")
	path := ctx.Request.URL.Path
	// 路径
	obj := strings.TrimPrefix(path, global.OMS_CONFIG.System.RouterPrefix)
	// 方法
	act := ctx.Request.Method
	e := service.ServiceGroupApp.System.CasbinService.Casbin()
	ok, err := e.Enforce(sub, dom, obj, act)
	if err != nil {
		response.Fail(nil, "权限获取失败:"+err.Error(), ctx)
		ctx.Abort()
		return
	}
	if !ok {
		response.Fail(nil, "权限不足", ctx)
		ctx.Abort()
		return
	}
	ctx.Next()
}
