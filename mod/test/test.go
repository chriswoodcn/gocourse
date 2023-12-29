package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	url := "www.taobao.com"
	defer logger.Sync() // flushes buffer, if any
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
