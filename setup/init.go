package setup

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
	"go.uber.org/zap"
)

func Setup() {
	// 初始化配置文件（首先）
	InitViper()
	// 初始化Logger
	global.Logger = initZap()
	// 初始化DB
	global.DB = initDB()
	// 初始化Redis
	global.Redis = initRedis()
	// 初始化casbin
	global.Enforcer = initEnforcer()

	// setup success info
	global.Logger.Info("success setup",
		zap.String("global.Mode", helper.GetMode()))
}
