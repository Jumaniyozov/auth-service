package main

import (
	"github.com/jumaniyozov/auth-service/internal/app"
	"github.com/jumaniyozov/auth-service/internal/config"
	"github.com/jumaniyozov/auth-service/internal/lib/setups"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := setups.SetupLogger(cfg.Env)

	log.Info("starting application", slog.Int("port", cfg.GRPC.Port))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	receivedSignal := <-stop
	log.Info("received signal", slog.String("signal", receivedSignal.String()))

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
