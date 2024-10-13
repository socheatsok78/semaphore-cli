package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Version is the version of the CLI
	Version = "dev"

	// Command-line root command
	rootCmd = &cobra.Command{
		Use:     "semaphore",
		Short:   "A backup and restore tool for Semaphore CI",
		Version: Version,
	}
)

var (
	configDNS      string
	configAddr     string
	configUsername string
	configPassword string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configDNS, "dns", "127.0.0.11", "dns resolver")
	rootCmd.PersistentFlags().StringVar(&configAddr, "addr", "https://cloud.semaphoreui.com", "the semaphore address")
	rootCmd.PersistentFlags().StringVar(&configUsername, "username", "", "the semaphore username")
	rootCmd.PersistentFlags().StringVar(&configPassword, "password", "", "the semaphore password")
}

func Execute() error {
	return rootCmd.Execute()
}
