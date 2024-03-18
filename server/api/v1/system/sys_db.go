package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/service"
)

type DbApi struct{}

func (*DbApi) CheckDB(c *gin.Context) {
	if global.OMS_DB != nil {
		response.Success(gin.H{"needInit": false}, "数据库无需初始化", c)
	} else {
		response.Warn(gin.H{"needInit": true}, "数据库尚未初始化", c)
	}
}

func (*DbApi) InitDB(c *gin.Context) {
	if global.OMS_DB != nil {
		response.Warn(nil, "数据库无需初始化", c)
		return
	}

	if err := service.ServiceGroupApp.System.DbService.InitDB(); err != nil {
		global.OMS_LOG.Error("初始化数据库失败")
		response.Fail(nil, "初始化数据库失败", c)
		return
	}

	response.Success(nil, "初始化数据库成功", c)
}

func (*DbApi) UpdateDB(c *gin.Context) {
	global.OMS_LOG.Info("升级数据库成功")
	global.OMS_LOG.Error("升级数据库失败")
}
