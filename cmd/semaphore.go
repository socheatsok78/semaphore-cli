package cmd

import "github.com/spf13/cobra"

var (
	// Version is the version of the CLI
	Version = "dev"

	// Command-line root command
	rootCmd = &cobra.Command{
		Use:     "semaphore",
		Short:   "A backup and restore tool for Semaphore CI",
		Version: Version,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

var (
	dnsResolver       string
	semaphoreURL      string
	semaphoreUsername string
	semaphorePassword string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&dnsResolver, "dns", "127.0.0.11", "DNS Resolver")
	rootCmd.PersistentFlags().StringVar(&semaphoreURL, "semaphore", "https://cloud.semaphoreui.com", "Semaphore URL")
	rootCmd.PersistentFlags().StringVar(&semaphoreUsername, "semaphore-username", "", "Semaphore Username")
	rootCmd.PersistentFlags().StringVar(&semaphorePassword, "semaphore-password", "", "Semaphore Password")
}

func Execute() error {
	return rootCmd.Execute()
}
