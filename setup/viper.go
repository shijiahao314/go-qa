package setup

import (
	"os"
	"path/filepath"
	"runtime"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
	"github.com/spf13/viper"
)

const (
	DEV_CONFIG_FILE  = "config.dev.yaml"
	PROD_CONFIG_FILE = "config.prod.yaml"
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
	mode := helper.GetMode()

	v := viper.New()
	configFile := DEV_CONFIG_FILE

	switch mode {
	case global.TEST:
		gin.SetMode(gin.TestMode)
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(filepath.Dir(b))
		configFile = filepath.Join(path, DEV_CONFIG_FILE)
	case global.DEV:
		gin.SetMode(gin.DebugMode)
		configFile = DEV_CONFIG_FILE
	case global.PROD:
		gin.SetMode(gin.ReleaseMode)
		configFile = PROD_CONFIG_FILE
	}

	slog.Info("config file", slog.String("path", configFile), slog.String("mode", mode))

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
