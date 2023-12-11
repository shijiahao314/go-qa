package setup

import (
	"os"

	"github.com/shijiahao314/go-qa/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func initZap() *zap.Logger {
	var encoderConfig zapcore.EncoderConfig
	switch global.Mode {
	case global.DEV:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	case global.PROD:
		encoderConfig = zap.NewProductionEncoderConfig()
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	// output1: file
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
	// output2: console
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zapcore.DebugLevel,
	)
	logger := zap.New(zapcore.NewTee(fileCore, consoleCore))
	// logger, err := zap.NewDevelopment(zap.AddCaller())
	// if err != nil {
	// 	panic(err)
	// }
	return logger
}
