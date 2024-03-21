package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
)

type DbApi struct{}

func (*DbApi) CheckDB(c *gin.Context) {
	if err := systemDBService.CheckDB(); err != nil {
		fmt.Println("[Golang] 数据库尚未初始化: " + err.Error())
		response.Warn(gin.H{"needInit": true}, "数据库尚未初始化", c)
	} else {
		response.Success(gin.H{"needInit": false}, "数据库无需初始化", c)
	}
}

func (*DbApi) InitDB(c *gin.Context) {
	if err := systemDBService.CheckDB(); err != nil {
		fmt.Println("[Golang] 数据库需要初始化: " + err.Error())
		if err := initDBService.InitDB(); err != nil {
			global.OMS_LOG.Error("[Golang] 初始化数据库失败" + ": " + err.Error())
			response.Fail(nil, "初始化数据库失败", c)
		} else {
			fmt.Println("[Golang] 初始化数据库成功")
			response.Success(nil, "初始化数据库成功", c)
		}
	} else {
		fmt.Println("[Golang] 数据库无需初始化")
		response.Warn(nil, "数据库无需初始化", c)
	}
}

func (*DbApi) UpdateDB(c *gin.Context) {
	global.OMS_LOG.Info("[Golang] 升级数据库成功")
	global.OMS_LOG.Error("[Golang] 升级数据库失败")
}
