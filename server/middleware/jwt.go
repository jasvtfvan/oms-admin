package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
	jwtFreecache "github.com/jasvtfvan/oms-admin/server/utils/freecache"
	jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"
)

// group白名单
var GroupWhiteList = []string{
	"/user/profile",
	"/update/check",
	"/update/db",
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			1.获取令牌
		*/
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.Unauthorized(nil, "未携带令牌，非法访问", ctx)
			return
		}
		// 判断组织编号空
		path := ctx.Request.URL.Path
		needGroup := !utils.Contains(GroupWhiteList, path) // 不在白名单
		var orgCode string
		if needGroup {
			orgCode = ctx.Request.Header.Get("x-group")
			if orgCode == "" {
				response.BadReq(nil, "组织编号不能为空", ctx)
				return
			}
		}
		/*
			2.验证令牌
		*/
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.Unauthorized(nil, "令牌已过期，需重新登录", ctx)
				return
			}
			response.Unauthorized(nil, err.Error(), ctx)
			return
		}
		// 用户禁用/删除的逻辑，不在此验证。用户禁用/删除，执行踢人逻辑，使得redis中没有缓存
		/*
			3.通过缓存比对令牌
		*/
		username := claims.Username

		if global.OMS_CONFIG.System.AuthCache == "redis" {
			// 不能放在声明区，因为拿不到global
			var jwtStore = jwtRedis.GetRedisStore().UseWithCtx(ctx)
			var cacheToken string = jwtStore.Get(username, false)
			if token != cacheToken {
				response.Unauthorized(nil, "其他客户端登录，令牌已失效", ctx)
				return
			}
		} else {
			// 不能放在声明区，因为拿不到global
			var jwtStore = jwtFreecache.GetStoreJWT().UseWithCtx(ctx)
			var cacheToken string = jwtStore.Get(username, false)
			if token != cacheToken {
				response.Unauthorized(nil, "其他客户端登录，令牌已失效", ctx)
				return
			}
		}
		/*
			4.判断用户是否具有x-group权限
		*/
		if needGroup {
			groups := claims.Groups
			if !utils.Contains(groups, orgCode) {
				response.BadReq(nil, "没有组织权限", ctx)
				return
			}
		}
		ctx.Set("claims", claims)
		ctx.Next()
		/*
			5.换新token，过期时间 - 现在时间 < 缓冲时间，就需要换token
		*/
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			exp, _ := utils.ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(exp)) // 过期时间
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			ctx.Header("new-token", newToken)
			if global.OMS_CONFIG.System.AuthCache == "redis" {
				var jwtStore = jwtRedis.GetRedisStore().UseWithCtx(ctx)
				jwtStore.Set(username, newToken)
			} else {
				var jwtStore = jwtFreecache.GetStoreJWT().UseWithCtx(ctx)
				jwtStore.Set(username, newToken)
			}
		}
	}
}
