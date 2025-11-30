package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/config"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/infra/migrations"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/infra/postgresql"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "migrateUp",
	Short: "It migrates up user-service database schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateUp called")

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

		migrator, err := migrations.NewMigrate(postgresqlDB, postgresql.Name)
		if err != nil {
			logger.Error("failed to create migrator", "error", err)
			os.Exit(1)
		}

		if err := migrator.UP(); err != nil {
			logger.Error("failed to migrate", "error", err)
			os.Exit(1)
		}

		defer func() {
			if err := migrator.Close(); err != nil {
				logger.Error("failed to close migrator", "error", err)
				os.Exit(1)
			}
		}()
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
}
