package initialize

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/middleware"
	"github.com/jasvtfvan/oms-admin/server/router"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {

	switch global.OMS_CONFIG.System.Env {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	if global.OMS_CONFIG.System.Env != "release" {
		// gin.Logger默认控制台输出请求信息：时间、地址、ip、返回状态、响应时间
		r.Use(gin.Logger())
	}

	// 禁止访问文件目录，只保留单个文件的访问权限
	r.StaticFS(global.OMS_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.OMS_CONFIG.Local.StorePath)})
	if global.OMS_CONFIG.System.UseTls {
		// 使用中间件加载https
		r.Use(middleware.LoadTls())
	}
	if global.OMS_CONFIG.System.Env == "release" {
		r.Use(middleware.CorsByRules())
	} else {
		r.Use(middleware.Cors())
	}

	// 注册路由-不做鉴权
	publicGroup := r.Group(global.OMS_CONFIG.System.RouterPrefix)
	{
		// 健康检测
		publicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
	{
		router.InitPublicRouter(publicGroup)
	}

	// 注册路由-鉴权
	privateGroup := r.Group(global.OMS_CONFIG.System.RouterPrefix)
	// privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitPrivateRouter(privateGroup)
	}

	global.OMS_LOG.Info("router register success! 路由注册成功!")
	return r
}
