package system

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	sysRes "github.com/jasvtfvan/oms-admin/server/model/system/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"github.com/jasvtfvan/oms-admin/server/utils/crypto"
)

type DbApi struct{}

// 这里是单机，如果是集群则需要分布式锁，比如使用redis
var (
	dbOnce sync.Once
)

// CheckInit
// @Tags	db
// @Summary	检查DB是否初始化
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}	"返回提示信息"
// @Router	/init/check [post]
func (*DbApi) CheckInit(c *gin.Context) {
	if err := initDBService.CheckDB(); err != nil {
		fmt.Println("[Golang] DB尚未初始化: " + err.Error())
		response.Fail(gin.H{"ready": false}, "DB尚未初始化", c)
	} else {
		initDBService.ClearInitializer()
		response.Success(gin.H{"ready": true}, "DB已准备就绪", c)
	}
}

// InitDB
// @Tags	db
// @Summary	初始化DB
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}	"返回提示信息"
// @Router	/init/db [post]
func (*DbApi) InitDB(c *gin.Context) {
	/*
		先判断初始化密码是否正确
	*/
	var isParamsOk bool = true
	params, err := c.GetRawData()
	if err != nil {
		isParamsOk = false
	}
	// 反序列化入参{"secret": "xxxxxxxxxx"}
	var req map[string]interface{}
	if isParamsOk {
		err = json.Unmarshal(params, &req) // 反序列化
		if err != nil {
			isParamsOk = false
		}
	}
	// 拿到secret值
	var secret string
	if isParamsOk {
		var ok bool
		secret, ok = req["secret"].(string)
		if !ok {
			isParamsOk = false
		}
	}
	// 根据secret反序列化{"initPwd":"Oms123Admin456"}
	var jsonObj map[string]interface{}
	if isParamsOk {
		secretDecrypted := crypto.AesDecrypt(secret)            // 对称解密
		err = json.Unmarshal([]byte(secretDecrypted), &jsonObj) // 反序列化
		if err != nil {
			isParamsOk = false
		}
	}
	// 拿到initPwd值
	var initPwdStr string
	if isParamsOk {
		var ok bool
		initPwdStr, ok = jsonObj["initPwd"].(string)
		if !ok {
			isParamsOk = false
		}
	}
	// 判断initPwd和系统配置的密码是否相等
	if isParamsOk {
		initPwd := global.OMS_CONFIG.System.InitPwd
		if initPwd != initPwdStr {
			response.Fail(nil, "初始化密码错误", c)
			return
		} // 不需要else，如果没有错误，就向下执行了
	}
	if !isParamsOk {
		response.Fail(nil, "请求参数错误", c)
		return
	}

	if err := initDBService.CheckDB(); err != nil {
		var firstGoroutine bool = false
		// 这里是单机，如果是集群则需要分布式锁
		dbOnce.Do(func() {
			fmt.Println("[Golang] DB尚未初始化: " + err.Error())
			if err := initDBService.InitDB(); err != nil {
				// 初始化失败，写入fatal日志，因为代码错误，会导致程序不能正确运行
				global.OMS_LOG.Fatal("初始化DB失败" + ": " + err.Error())
				response.Fail(nil, "初始化DB失败，检查服务后重启", c)
			} else {
				// 初始化时清除缓存
				global.OMS_REDIS.Clear(c)
				global.OMS_FREECACHE.Clear()
				fmt.Println("[Golang] 初始化DB成功")
				response.Success(nil, "初始化DB成功", c)
			}
			firstGoroutine = true
		})
		if !firstGoroutine {
			response.Fail(gin.H{"needRefresh": true}, "初始化进行中", c)
		}
	} else {
		fmt.Println("[Golang] DB已准备就绪")
		response.Success(nil, "DB已准备就绪", c)
	}
}

// CheckUpdate
// @Tags	db
// @Summary	检查更新
// @Security  ApiKeyAuth
// @Security  ApiKeyDomain
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=sysRes.SysDB,msg=string}	"返回提示信息"
// @Router	/update/check [post]
func (*DbApi) CheckUpdate(c *gin.Context) {
	rootUsername := global.OMS_CONFIG.System.Username
	v := c.Value("claims")
	if claims, ok := v.(*utils.CustomClaims); ok {
		if claims.Username != rootUsername {
			response.Fail(sysRes.SysDB{
				Updated:    false,
				OldVersion: "",
				NewVersion: "",
			}, "系统超级管理员才有权限", c)
			return
		}
	} else {
		response.Fail(sysRes.SysDB{
			Updated:    false,
			OldVersion: "",
			NewVersion: "",
		}, "解析令牌信息失败", c)
		return
	}
	if v1, v2, err := updateDBService.CheckUpdate(); err != nil {
		fmt.Println("[Golang] DB查询失败: " + err.Error())
		response.Fail(sysRes.SysDB{
			Updated:    false,
			OldVersion: "",
			NewVersion: "",
		}, "DB查询失败", c)
	} else {
		if v1 == v2 {
			updateDBService.ClearUpdater()
			response.Success(sysRes.SysDB{
				Updated:    true,
				OldVersion: v1,
				NewVersion: v2,
			}, "DB已升级", c)
		} else {
			fmt.Println(fmt.Sprintf("[Golang] DB需要升级: %s -> %s", v1, v2))
			response.Fail(sysRes.SysDB{
				Updated:    false,
				OldVersion: v1,
				NewVersion: v2,
			}, "DB需要升级", c)
		}
	}
}

// UpdateDB
// @Tags	db
// @Summary	升级DB
// @Security  ApiKeyAuth
// @Security  ApiKeyDomain
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=sysRes.SysDB,msg=string}	"返回提示信息"
// @Router	/update/db [post]
func (*DbApi) UpdateDB(c *gin.Context) {
	rootUsername := global.OMS_CONFIG.System.Username
	v := c.Value("claims")
	if claims, ok := v.(*utils.CustomClaims); ok {
		if claims.Username != rootUsername {
			response.Fail(sysRes.SysDB{
				Updated:    false,
				OldVersion: "",
				NewVersion: "",
			}, "系统超级管理员才有权限", c)
			return
		}
	} else {
		response.Fail(sysRes.SysDB{
			Updated:    false,
			OldVersion: "",
			NewVersion: "",
		}, "解析令牌信息失败", c)
		return
	}
	if v1, v2, err := updateDBService.CheckUpdate(); err != nil {
		fmt.Println("[Golang] DB查询失败: " + err.Error())
		response.Fail(nil, "DB查询失败", c)
	} else {
		if v1 == v2 {
			updateDBService.ClearUpdater()
			response.Success(gin.H{"updated": true, "oldVersion": v1, "newVersion": v2}, "DB已升级", c)
		} else {
			fmt.Println(fmt.Sprintf("[Golang] DB需要升级: %s -> %s", v1, v2))
			if err := updateDBService.UpdateDB(); err != nil {
				global.OMS_LOG.Error("[Golang] 升级DB失败" + ": " + err.Error())
				response.Fail(sysRes.SysDB{
					Updated:    false,
					OldVersion: v1,
					NewVersion: v2,
				}, "升级DB失败", c)
			} else {
				// 初始化时清除缓存
				global.OMS_REDIS.Clear(c)
				global.OMS_FREECACHE.Clear()
				fmt.Println("[Golang] 升级DB成功")
				response.Success(sysRes.SysDB{
					Updated:    true,
					OldVersion: v1,
					NewVersion: v2,
				}, "升级DB成功", c)
			}
		}
	}
}
