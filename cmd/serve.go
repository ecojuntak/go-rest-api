package cmd

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/server"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "To run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config()

		dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
		db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
		if err != nil {
			log.Panicf("cannot connect to database: %s\n", err.Error())
		}

		server := server.NewServer(config.Server.Host, config.Server.Port, db)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
