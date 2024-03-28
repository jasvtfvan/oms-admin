package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1.获取令牌
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.Fail(gin.H{"reload": true}, "未携带令牌，非法访问", ctx)
			ctx.Abort()
			return
		}
		// 2.验证令牌
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.Fail(gin.H{"reload": true}, "令牌已过期", ctx)
				ctx.Abort()
				return
			}
			response.Fail(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		// 用户禁用/删除的逻辑，不在此验证。用户禁用/删除，执行踢人逻辑，使得redis中没有缓存
		// 3.通过缓存比对令牌
		ctx.Set("claims", claims)
	}
}
