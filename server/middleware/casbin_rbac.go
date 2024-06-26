package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	sysRes "github.com/jasvtfvan/oms-admin/server/model/system/response"
	"github.com/jasvtfvan/oms-admin/server/service"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rootUsername := global.OMS_CONFIG.System.Username
		v := ctx.Value("claims")
		if claims, ok := v.(*utils.CustomClaims); ok { // 断言成功
			if claims.Username == rootUsername { // 超级管理员，所有接口都可以访问
				ctx.Next()
			} else { // 非超级管理员
				if isCasbinSource(ctx) { // 需要casbin验证的接口或资源
					handler(ctx, claims)
				} else { // 不需要验证的直接通过
					ctx.Next()
				}
			}
		} else {
			response.Unauthorized(nil, "解析令牌信息失败", ctx)
			return
		}
	}
}

// 判断是否为casbin需要验证的资源
func isCasbinSource(ctx *gin.Context) bool {
	// 路径
	path := ctx.Request.URL.Path
	obj := strings.TrimPrefix(path, global.OMS_CONFIG.System.RouterPrefix)
	// 方法
	act := ctx.Request.Method
	casbinInfos := sysRes.DefaultCasbinSource()
	for _, v := range casbinInfos {
		if v.Path == obj && v.Method == act {
			return true
		}
	}
	return false
}

func handler(ctx *gin.Context, claims *utils.CustomClaims) {
	// 角色
	roles := claims.Roles
	// 域
	dom := ctx.Request.Header.Get("x-group")
	if dom == "" {
		response.BadReq(nil, "组织编号不能为空", ctx)
		return
	}
	// 路径
	path := ctx.Request.URL.Path
	obj := strings.TrimPrefix(path, global.OMS_CONFIG.System.RouterPrefix)
	// 方法
	act := ctx.Request.Method
	e := service.ServiceGroupApp.System.CasbinService.Casbin()
	isOk := false
	for _, role := range roles {
		sub := role
		ok, _ := e.Enforce(sub, dom, obj, act)
		if ok {
			isOk = true
			break
		}
	}
	if isOk {
		ctx.Next()
		return
	} else {
		response.BadReq(nil, "没有该资源权限", ctx)
		return
	}
}
