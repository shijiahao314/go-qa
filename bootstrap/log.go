package bootstrap

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/shijiahao314/go-qa/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// mustInitZap 初始化 Logger
func mustInitZap() *zap.Logger {
	slog.Info("start to init logger")
	defer slog.Info("success init logger")

	var encoderConfig zapcore.EncoderConfig
	mode := global.Mode
	switch mode {
	case global.TEST:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	case global.DEV:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	case global.PROD:
		encoderConfig = zap.NewProductionEncoderConfig()
	default:
		slog.Error("invalid mode", slog.String("mode", string(mode)))
		panic(fmt.Sprintf("invalid mode: %s", mode))
	}
	// 日志输出流1: 日志文件
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/tmp.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}),
		zap.DebugLevel,
	)
	// 日志输出流2: 控制台
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zapcore.DebugLevel,
	)
	logger := zap.New(zapcore.NewTee(fileCore, consoleCore))

	return logger
}
