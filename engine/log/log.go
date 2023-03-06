package log

import (
	"dkz.com/engine/cfg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(cfg cfg.BaseConfig) *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writeSyncer := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   "log/foo.log",
			MaxSize:    10,
			MaxAge:     28,
			MaxBackups: 3,
			Compress:   false,
			LocalTime:  false,
		})
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	return zap.New(core)
}
