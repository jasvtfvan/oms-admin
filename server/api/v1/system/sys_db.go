package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
)

type DbApi struct{}

func (*DbApi) CheckInit(c *gin.Context) {
	if err := initDBService.CheckDB(); err != nil {
		fmt.Println("[Golang] DB尚未初始化: " + err.Error())
		response.Fail(gin.H{"ready": false}, "DB尚未初始化", c)
	} else {
		initDBService.ClearInitializer()
		response.Success(gin.H{"ready": true}, "DB已准备就绪", c)
	}
}

func (*DbApi) InitDB(c *gin.Context) {
	if err := initDBService.CheckDB(); err != nil {
		fmt.Println("[Golang] DB尚未初始化: " + err.Error())
		if err := initDBService.InitDB(); err != nil {
			global.OMS_LOG.Error("[Golang] 初始化DB失败" + ": " + err.Error())
			response.Fail(nil, "初始化DB失败", c)
		} else {
			fmt.Println("[Golang] 初始化DB成功")
			response.Success(nil, "初始化DB成功", c)
		}
	} else {
		fmt.Println("[Golang] DB已准备就绪")
		response.Success(nil, "DB已准备就绪", c)
	}
}

func (*DbApi) CheckUpdate(c *gin.Context) {
	if err := updateDBService.CheckUpdate(); err != nil {
		fmt.Println("[Golang] DB需要升级: " + err.Error())
		response.Fail(gin.H{"updated": false}, "DB需要升级", c)
	} else {
		updateDBService.ClearUpdater()
		response.Success(gin.H{"updated": true}, "DB已升级", c)
	}
}

func (*DbApi) UpdateDB(c *gin.Context) {
	if err := updateDBService.CheckUpdate(); err != nil {
		fmt.Println("[Golang] DB需要升级: " + err.Error())
		if err := updateDBService.UpdateDB(); err != nil {
			global.OMS_LOG.Error("[Golang] 升级DB失败" + ": " + err.Error())
			response.Fail(nil, "升级DB失败", c)
		} else {
			fmt.Println("[Golang] 升级DB成功")
			response.Success(nil, "升级DB成功", c)
		}
	} else {
		updateDBService.ClearUpdater()
		response.Success(nil, "DB已升级", c)
	}
}
