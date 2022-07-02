package cmd

import (
	"go-rest-api/config"
	"go-rest-api/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "To run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.ServerConfig()

		server := server.NewServer(config.Host, config.Port)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
