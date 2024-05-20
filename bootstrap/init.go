package bootstrap

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
	"go.uber.org/zap"
)

// 必须初始化
func MustInit() {
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
	// 初始化etcd
	global.Etcd = initEtcd()

	// setup success info
	global.Logger.Info("success setup",
		zap.String("global.Mode", string(helper.GetMode())))
}
