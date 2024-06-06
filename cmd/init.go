package cmd

import (
	"IOTProject/cmd/config"
	"IOTProject/cmd/create"
	"IOTProject/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "app",
	Short:        "app",
	SilenceUsage: true,
	Long:         `app`,
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(create.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
