package setup

import (
	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/helper"
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
