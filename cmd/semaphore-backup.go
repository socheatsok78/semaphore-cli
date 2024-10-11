package cmd

import (
	"github.com/socheatsok78/semaphore-cli/semaphore"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().String("project", "", "semaphore project id")
	backupCmd.MarkFlagRequired("project")
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup of Semaphore project",
	Run: func(cmd *cobra.Command, args []string) {
		httpClient := createHttpClient(cmd)
		s := semaphore.New(httpClient)
		projectID := cmd.Flag("project").Value.String()
		s.Backup(projectID)
	},
}
