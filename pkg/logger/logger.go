package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var Service = "logger"

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("service", Service)
	slog.SetDefault(logger)
}
