package initialize

import (
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func BaseInit() {
	_, err := utils.ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.OMS_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	rootUsername := global.OMS_CONFIG.System.Username
	if rootUsername == "" {
		panic(errors.New("系统管理员用户名不能为空"))
	}
	rootPassword := global.OMS_CONFIG.System.Password
	if rootPassword == "" {
		panic(errors.New("系统管理员密码不能为空"))
	}
}
