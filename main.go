package main

import (
	"fmt"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/setup"
	"go.uber.org/zap"
)

func main() {
	setup.Setup()

	r := setup.InitRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Port)); err != nil {
		global.Logger.Error("failed to run server", zap.String("host", "0.0.0.0"), zap.Int("port", global.Config.Port), zap.String("err", err.Error()))
	}
}
