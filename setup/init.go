package setup

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
	"go.uber.org/zap"
)

func Setup() {
	InitViper()
	global.DB = initDB()
	global.Logger = initZap()
	global.Redis = initRedis()

	global.Logger.Info("success setup",
		zap.String("global.Mode", helper.GetMode()))
}
