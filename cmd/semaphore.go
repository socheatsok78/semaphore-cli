package cmd

import (
	"github.com/spf13/cobra"
)

var (
	dnsResolverAddr   string
	semaphoreAddr     string
	semaphoreUsername string
	semaphorePassword string
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
	rootCmd.PersistentFlags().StringVar(&dnsResolverAddr, "dns", "127.0.0.11", "DNS Resolver")
	rootCmd.PersistentFlags().StringVar(&semaphoreAddr, "addr", "https://cloud.semaphoreui.com", "Semaphore URL")
	rootCmd.PersistentFlags().StringVar(&semaphoreUsername, "username", "", "Semaphore Username")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.PersistentFlags().StringVar(&semaphorePassword, "password", "", "Semaphore Password")
	rootCmd.MarkPersistentFlagRequired("password")
}

func Execute() error {
	return rootCmd.Execute()
}
