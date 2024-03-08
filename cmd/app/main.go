package main

import (
	"article-3-how-use-zerolog/config"
	"article-3-how-use-zerolog/pkg/logger"
	"errors"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.dev", ".env")
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		panic(err)
	}
	logger, err := logger.NewLoggerFromConfig(cfg.Logger)
	if err != nil {
		panic(err)
	}
	logger.Debug("Test..")
	logger.Error(errors.New("ASDSA"))
	appLog := logger.AddService("app")
	appLog.Info("ïnfo")

}
