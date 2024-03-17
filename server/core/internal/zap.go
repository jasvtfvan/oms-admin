package internal

import (
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

// GetEncoder 获取 zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoder(keys ...string) zapcore.Encoder {
	if global.OMS_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig(keys...))
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig(keys...))
}

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderConfig(keys ...string) zapcore.EncoderConfig {
	var config zapcore.EncoderConfig
	if len(keys) <= 0 {
		config = zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  global.OMS_CONFIG.Zap.StacktraceKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    global.OMS_CONFIG.Zap.ZapEncodeLevel(),
			EncodeTime:     z.CustomTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		}
		return config
	}

	config = zapcore.EncoderConfig{}
	if _slice(keys).includes("MessageKey") {
		config.MessageKey = "message"
	}
	if _slice(keys).includes("LevelKey") {
		config.LevelKey = "level"
	}
	if _slice(keys).includes("TimeKey") {
		config.TimeKey = "time"
	}
	if _slice(keys).includes("NameKey") {
		config.NameKey = "logger"
	}
	if _slice(keys).includes("CallerKey") {
		config.CallerKey = "caller"
	}
	if _slice(keys).includes("StacktraceKey") {
		config.StacktraceKey = global.OMS_CONFIG.Zap.StacktraceKey
	}
	if _slice(keys).includes("LineEnding") {
		config.LineEnding = zapcore.DefaultLineEnding
	}
	if _slice(keys).includes("EncodeLevel") {
		config.EncodeLevel = global.OMS_CONFIG.Zap.ZapEncodeLevel()
	}
	if _slice(keys).includes("EncodeTime") {
		config.EncodeTime = z.CustomTimeEncoder
	}
	if _slice(keys).includes("EncodeDuration") {
		config.EncodeDuration = zapcore.SecondsDurationEncoder
	}
	if _slice(keys).includes("EncodeCaller") {
		config.EncodeCaller = zapcore.FullCallerEncoder
	}

	return config
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc, keys ...string) zapcore.Core {
	writer := FileRotateLogs.GetWriteSyncer(l.String()) // 日志分割
	return zapcore.NewCore(z.GetEncoder(keys...), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.OMS_CONFIG.Zap.Prefix + " " + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetZapCores(keys ...string) []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.OMS_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level), keys...))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

type _slice []string

func (slice _slice) includes(str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
