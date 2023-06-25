package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("hoge", "fuga", "piyo")

	sl := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	sl.Info("hoge", "fuga", "piyo")

	sl.WithGroup("group1").Info("hoge", "fuga", "piyo")

	sl.With("key1", "value1").Info("hoge", "fuga", "piyo")
}
