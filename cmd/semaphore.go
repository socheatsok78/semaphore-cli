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

func init() {
	rootCmd.PersistentFlags().String("dns", "127.0.0.11", "dns resolver")
	rootCmd.PersistentFlags().String("addr", "https://cloud.semaphoreui.com", "the semaphore address")
	rootCmd.PersistentFlags().String("username", "", "the semaphore username")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.PersistentFlags().String("password", "", "the semaphore password")
	rootCmd.MarkPersistentFlagRequired("password")
}

func Execute() error {
	return rootCmd.Execute()
}
