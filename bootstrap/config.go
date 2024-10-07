package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/global"
	"github.com/spf13/viper"
)

const (
	DevConfigFile  = "config.dev.yaml"
	ProdConfigFile = "config.prod.yaml"
)

// mustInitViper 初始化配置文件
func mustInitViper() {
	slog.Info("start to init config file")
	defer slog.Info("success init config file")

	mode := mustInitMode()
	global.Mode = mode

	v := viper.New()
	configFile := DevConfigFile

	switch mode {
	case global.TEST:
		gin.SetMode(gin.TestMode)
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(filepath.Dir(b))
		configFile = filepath.Join(path, DevConfigFile)
	case global.DEV:
		gin.SetMode(gin.DebugMode)
		configFile = DevConfigFile
	case global.PROD:
		gin.SetMode(gin.ReleaseMode)
		configFile = ProdConfigFile
	default:
		slog.Error("invalid mode", slog.String("mode", string(mode)))
		panic(fmt.Sprintf("invalid mode: %s", mode))
	}

	// global.Logger 还未初始化，使用 slog
	slog.Info("use config file", slog.String("path", configFile), slog.String("mode", string(mode)))

	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}

// mustInitMode 初始化运行模式
func mustInitMode() (mode global.AppMode) {
	slog.Info("start to init mode")
	defer func() {
		slog.Info("success init mode", slog.String("mode", string(mode)))
	}()

	envMode, ok := os.LookupEnv("MODE")
	if ok {
		switch envMode {
		case string(global.TEST):
			mode = global.TEST
		case string(global.DEV):
			mode = global.DEV
		case string(global.PROD):
			mode = global.PROD
		default:
			mode = global.DefaultAppMode
		}
	} else {
		mode = global.DefaultAppMode
	}

	return
}
