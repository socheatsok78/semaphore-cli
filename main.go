package main

import (
	"os"

	"github.com/go-kit/log/level"
	"github.com/socheatsok78/semaphore-cli/cmd"
	"github.com/socheatsok78/semaphore-cli/internal"
)

func main() {
	if err := cmd.Execute(); err != nil {
		level.Error(internal.Logger).Log("msg", "Failed to execute command", "err", err)
		os.Exit(1)
	}
}
