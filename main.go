package main

import (
	"fmt"
	"log/slog"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/setup"
)

func main() {
	setup.Setup()

	r := setup.InitRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Port)); err != nil {
		slog.Error("failed to run server on 0.0.0.0", slog.Int("port", global.Config.Port), err)
	}
}
