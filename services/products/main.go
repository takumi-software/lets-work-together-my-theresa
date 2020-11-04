package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/takumi-software/lets-work-together-my-theresa/pkg/app"
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal"
	"go.uber.org/zap"
)

func main() {
	var cfg internal.Config

	command, logger, ctx := app.NewCobraApp(internal.Name, &cfg)
	command.SilenceUsage = true
	command.RunE = func(*cobra.Command, []string) error {
		return internal.Bootstrap(ctx, cfg, logger)
	}

	if err := command.Execute(); err != nil {
		logger.Error(
			"unable to start backend service",
			zap.Error(err),
			zap.Stack("stacktrace"),
		)
		os.Exit(1)
	}
}
