package setup

import (
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

func InitMode() {
	mode, ok := os.LookupEnv("MODE")
	if ok {
		switch mode {
		case global.TEST:
			global.Mode = global.TEST
		case global.DEV:
			global.Mode = global.DEV
		case global.PROD:
			global.Mode = global.PROD
		}
	} else {
		global.Mode = global.DEFAULT_MODE
	}
}

func InitViper() {
	InitMode()
	mode := global.Mode

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
	}

	slog.Info("config file", slog.String("path", configFile), slog.String("mode", string(mode)))

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
