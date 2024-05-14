package main

import (
	"context"
	"flag"
	"github.com/archdx/zcfg"
	"github.com/rs/zerolog"
	"github.com/sanches1984/referral-bot/config"
	"github.com/sanches1984/referral-bot/internal/app/service"
	"github.com/sanches1984/referral-bot/internal/app/storage"
	stdLog "log"
	"os"
	"os/signal"
	"syscall"
)

var version = "devel"

func main() {
	cfg := config.NewDefaultConfig()
	err := zcfg.New(cfg, zcfg.FromFile("config.yaml"), zcfg.UseFlags(flag.CommandLine)).Load()
	if err != nil {
		stdLog.Fatalf("load config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		flag.Usage()
		stdLog.Fatalf("validate config: %v", err)
	}

	ctx, stop := context.WithCancel(context.Background())
	logger := zerolog.New(os.Stderr).Level(cfg.LogLevel).With().Timestamp().Str("ver", version).Logger()

	db, closer, err := storage.New(ctx, "db_credentials.json")
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize storage")
		return
	}
	defer closer()

	srv, err := service.New(cfg.Bot.Token, db, logger)
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize service")
		return
	}

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		logger.Info().Msg("termination signal received")
		stop()
	}()

	srv.Run(ctx)
}
