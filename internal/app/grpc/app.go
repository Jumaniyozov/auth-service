package grpcapp

import (
	"fmt"
	authgrpc "github.com/jumaniyozov/auth-service/internal/grpc/auth"
	authSrvc "github.com/jumaniyozov/auth-service/internal/services/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(logger *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	auth := authSrvc.New(
		logger,
		nil,
		nil,
		nil,
		0)

	authgrpc.Register(gRPCServer, auth)

	return &App{
		log:        logger,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc sering is running", slog.String("address", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping grpc server")

	a.gRPCServer.GracefulStop()
}
