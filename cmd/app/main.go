package main

import (
	_ "github.com/manzanit0/logger-idea/pkg/logger"
	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("test log")
}
