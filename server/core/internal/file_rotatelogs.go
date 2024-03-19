package internal

import (
	"os"

	"github.com/jasvtfvan/oms-admin/server/global"
	"go.uber.org/zap/zapcore"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer

func (r *fileRotateLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.OMS_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.OMS_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
