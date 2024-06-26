package core

import (
	"fmt"
	"os"

	"github.com/jasvtfvan/oms-admin/server/core/internal"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.OMS_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.OMS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.OMS_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.OMS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func ZapGin() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.OMS_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.OMS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.OMS_CONFIG.Zap.Director, os.ModePerm)
	}

	configKeys := []string{"MessageKey"}

	cores := internal.Zap.GetZapCores(configKeys...)
	logger = zap.New(zapcore.NewTee(cores...))

	return logger
}
