package setups

import (
	"github.com/jumaniyozov/auth-service/internal/constants/environment"
	"github.com/jumaniyozov/auth-service/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case environment.EnvLocal:
		log = SetupPrettySlog()
	case environment.EnvDev:
		log = SetupPrettySlog()
	case environment.EnvProd:
		log = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	default:
		panic("unknown environment")
	}

	return log
}

func SetupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
