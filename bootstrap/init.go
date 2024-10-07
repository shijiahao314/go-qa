package bootstrap

import (
	"log/slog"

	"github.com/shijiahao314/go-qa/global"
)

// MustInit 初始化必要配置，如果失败将 panic
func MustInit() {
	slog.Info("start to setup must init")
	defer slog.Info("success setup must init")

	// 初始化配置文件（首先）
	mustInitViper()
	// 初始化 Logger
	global.Logger = mustInitZap()
	// 初始化 DB
	global.DB = mustInitDB()
	// 初始化 Casbin
	global.Enforcer = mustInitCasbin()
}

// Init 初始化可选配置
func Init() {
	slog.Info("start to setup init")
	defer slog.Info("finish setup init")

	// 初始化 Redis
	redis, err := initRedis()
	if err != nil {
		global.Redis = redis
	}
}
