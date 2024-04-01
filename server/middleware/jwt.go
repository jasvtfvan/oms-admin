package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
	jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"
)

var jwtStore = jwtRedis.GetRedisStore()

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtStore = jwtStore.UseWithCtx(ctx)
		/*
			1.获取令牌
		*/
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.Fail(gin.H{"reload": true}, "未携带令牌，非法访问", ctx)
			ctx.Abort()
			return
		}
		/*
			2.验证令牌
		*/
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.Fail(gin.H{"reload": true}, "令牌已过期，需重新登录", ctx)
				ctx.Abort()
				return
			}
			response.Fail(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		// 用户禁用/删除的逻辑，不在此验证。用户禁用/删除，执行踢人逻辑，使得redis中没有缓存
		/*
			3.通过缓存比对令牌
		*/
		username := claims.Username
		var cacheToken string = jwtStore.Get(username, false)
		if token != cacheToken {
			response.Fail(gin.H{"reload": true}, "其他客户端登录，令牌已失效", ctx)
			ctx.Abort()
			return
		}
		/*
			过期时间 - 现在时间 < 缓冲时间，就需要换token
		*/
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			exp, _ := utils.ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(exp)) // 过期时间
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			ctx.Header("new-token", newToken)
			jwtStore.Set(username, newToken)
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
