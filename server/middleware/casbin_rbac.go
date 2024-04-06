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
	roles := claims.Roles
	// 域
	dom := ctx.Request.Header.Get("x-group")
	// 路径
	path := ctx.Request.URL.Path
	obj := strings.TrimPrefix(path, global.OMS_CONFIG.System.RouterPrefix)
	// 方法
	act := ctx.Request.Method
	e := service.ServiceGroupApp.System.CasbinService.Casbin()
	isOk := false
	for _, role := range roles {
		sub := role.RoleCode
		ok, _ := e.Enforce(sub, dom, obj, act)
		if ok {
			isOk = true
			break
		}
	}
	if isOk {
		ctx.Next()
	} else {
		response.Fail(nil, "权限不足", ctx)
		ctx.Abort()
		return
	}
}
