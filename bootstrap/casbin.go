package bootstrap

import (
	"log/slog"

	"github.com/casbin/casbin/v2"
)

// mustInitCasbin 初始化 casbin
func mustInitCasbin() *casbin.Enforcer {
	slog.Info("start to init casbin")
	defer slog.Info("success init casbin")

	return initEnforcer()
}
