package main

import (
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	suger := logger.Sugar()

	suger.Infow("failed to fetch URL",
		"url", "pornhub.com",
		"attempts", 5,
		"duration", time.Second*50)
}
