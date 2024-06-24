package main

import (
	"DevKit-Neuro-hub/internal/config"
	main_logger "DevKit-Neuro-hub/internal/logger"
	"DevKit-Neuro-hub/internal/server"
	main_validator "DevKit-Neuro-hub/internal/validator"
	"context"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer ctxCancel()

	mainValidator, err := main_validator.NewValidator()

	mainConfig, err := config.LoadConfig(mainValidator)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := main_logger.NewLogger(config.DebugLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("config loaded", zap.Any("config", mainConfig))

	mainServer, err := server.NewServer(mainConfig, logger, mainValidator)
	if err != nil {
		logger.Fatal("error while starting server", zap.Error(err))
	}

	defer mainServer.Conn.Close()

	mainServer.Run(ctx)
}
