package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/config"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/infra/postgresql"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/routes"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/middleware"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/server"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "It runs the user service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		cfg, err := config.GetInstance()
		if err != nil {
			logger.Error("failed to get config", "error", err)
			os.Exit(1)
		}

		postgresql := postgresql.NewPostgresql(
			postgresql.WithHost(cfg.Postgresql.Host),
			postgresql.WithPort(cfg.Postgresql.Port),
			postgresql.WithUser(cfg.Postgresql.User),
			postgresql.WithPassword(cfg.Postgresql.Password),
			postgresql.WithName(cfg.Postgresql.Name),
			postgresql.WithMaxOpenConn(cfg.Postgresql.MaxOpenConn),
			postgresql.WithMaxIdleConn(cfg.Postgresql.MaxIdleConn),
			postgresql.WithMaxIdleTime(cfg.Postgresql.MaxIdleTime),
			postgresql.WithSSLMode(cfg.Postgresql.SSLMode),
			postgresql.WithTimeout(cfg.Postgresql.Timeout),
		)

		postgresqlDB, err := postgresql.Connect()
		if err != nil {
			logger.Error("failed to connect to postgresql", "error", err)
			os.Exit(1)
		}

		defer func() {
			if err := postgresqlDB.Close(); err != nil {
				logger.Error("failed to close postgresql", "error", err)
				os.Exit(1)
			}
		}()

		apiError := helper.NewAPIError(logger)
		middlewares := middleware.NewMiddleware(cfg, apiError)
		healthHandler := handlers.NewHealthHandler(logger, cfg, apiError)
		healthRoute := routes.NewHealthRoute(healthHandler)
		registerRoutes := routes.NewRegister(
			routes.WithAPIError(apiError),
			routes.WithMiddleware(middlewares),
			routes.WithHealthRoute(healthRoute),
		)

		httpServer := server.NewServer(
			server.WithPort(cfg.Server.Port),
			server.WithHost(cfg.Server.Host),
			server.WithHandler(registerRoutes.RegisterRoutes()),
			server.WithIdleTimeout(cfg.Server.IdleTimeout),
			server.WithReadTimeout(cfg.Server.ReadTimeout),
			server.WithWriteTimeout(cfg.Server.WriteTimeout),
			server.WithErrorLog(slog.NewLogLogger(logger.Handler(), slog.LevelError)),
		)

		logger.Info("starting server", "addr", cfg.Server.Host+":"+cfg.Server.Port, "env", cfg.Application.Environment)

		if err := httpServer.Connect(); err != nil {
			logger.Error("failed to start http server", "error", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
