package main

import (
	"fmt"

	"github.com/shijiahao314/go-qa/bootstrap"
	"github.com/shijiahao314/go-qa/global"
	"go.uber.org/zap"
)

func main() {
	bootstrap.MustInit()

	r := bootstrap.InitRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Port)); err != nil {
		global.Logger.Error("failed to run server",
			zap.String("host", "0.0.0.0"),
			zap.Int("port", global.Config.Port),
			zap.String("err", err.Error()))
	}
}
