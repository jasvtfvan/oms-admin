package initialize

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/middleware"
	"github.com/jasvtfvan/oms-admin/server/router"
	"go.uber.org/zap"
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

func Routers(logger *zap.Logger) *gin.Engine {

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

	if logger != nil {
		// zap打印请求信息，同时打印到console和log文件
		r.Use(middleware.ZapLogger(logger))
	} else {
		// gin.Logger默认控制台输出请求信息：时间、返回状态、响应时间、ip、请求方法（GET、POST等）、请求相对路径
		// r.Use(gin.Logger())
		// 控制台打印请求信息
		r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// 自定义格式
			return fmt.Sprintf("[gin.LoggerWithFormatter] %s [%s] | \t %s %s [%s] | \t %d %s [%s] \n[gin.LoggerWithFormatter:UserAgent] %s \n",
				param.TimeStamp.Format(time.DateTime),
				param.ClientIP,
				param.Request.Proto,
				param.Method,
				param.Path,
				param.StatusCode,
				param.ErrorMessage,
				param.Latency.String(),
				param.Request.UserAgent(),
			)
		}))
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
