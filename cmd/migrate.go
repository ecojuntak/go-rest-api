package cmd

import (
	"fmt"
	"go-rest-api/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "To run database migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := config.DatabaseConfig()

		dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Name)
		m, err := migrate.New("file://./migrations", dbConnection)
		if err != nil {
			log.Fatalf("error creating database migrator: %s", err)
			return err
		}

		err = m.Up()
		if err == migrate.ErrNoChange {
			log.Println("no database migration changes")
			return nil
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
