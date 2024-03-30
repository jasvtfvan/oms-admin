package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rootUsername := global.OMS_CONFIG.System.Username
		v := ctx.Value("claims")
		if claims, ok := v.(utils.CustomClaims); ok {
			if claims.Username == rootUsername { // 超级管理员，所有接口都可以访问
				ctx.Next()
			} else {
				handler(ctx)
			}
		} else {
			response.Fail(nil, "解析令牌信息失败", ctx)
			ctx.Abort()
			return
		}
	}
}

func handler(ctx *gin.Context) {
	// 是否有权限，取决于角色list，并满足:
	// 1、接口属于casbin-api列表
	// 2、角色list中具有满足casbin的api权限
	// 3、角色list中具有所在群组的权限（或当前群组向上能追溯到list中的角色）
	group := ctx.Request.Header.Get("x-group")
	if group == "" {
		response.Fail(nil, "组织编号不能为空", ctx)
		ctx.Abort()
		return
	}
	ctx.Next()
}
