package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
)

type DbApi struct{}

func (*DbApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.OMS_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.OMS_LOG.Info(message)
	response.Success(gin.H{"needInit": needInit}, message, c)
}

func (*DbApi) InitDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.OMS_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.OMS_LOG.Info(message)
	response.Success(gin.H{"needInit": needInit}, message, c)
}

func (*DbApi) UpdateDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.OMS_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.OMS_LOG.Info(message)
	response.Success(gin.H{"needInit": needInit}, message, c)
}
