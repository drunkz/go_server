package log

import (
	"os"
	"time"

	"dkz.com/engine/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const LOGTIMEFORMAT = "2006-01-02 15:04:05.000"

func InitLog(cfg config.BaseConfig) *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder	// Windows Cmd环境颜色显示异常
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(LOGTIMEFORMAT))
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writeSyncer := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   cfg.LogFileName,
			MaxSize:    int(cfg.LogMaxSize),    // 单日志文件允许大小(MB)
			MaxAge:     int(cfg.LogMaxAge),     // 最大保留天数
			MaxBackups: int(cfg.LogMaxBackups), // 最大保留数量
			Compress:   cfg.LogCompress,        // 是否压缩保留日志
			LocalTime:  true,
		})
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, multiWriteSyncer, zapcore.Level(cfg.LogLevel))
	return zap.New(core, zap.AddCaller())
}
