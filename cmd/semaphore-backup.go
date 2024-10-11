package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().String("project", "", "Project ID")
	backupCmd.MarkFlagRequired("project")
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup of Semaphore project",
	Run: func(cmd *cobra.Command, args []string) {
		backup(cmd)
	},
}

func backup(cmd *cobra.Command) {}
