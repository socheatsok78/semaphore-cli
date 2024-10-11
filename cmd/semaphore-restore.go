package cmd

import (
	"github.com/socheatsok78/semaphore-cli/semaphore"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().String("project", "", "semaphore project id")
	restoreCmd.Flags().String("backup-file", "", "semaphore backup file to restore from")
	restoreCmd.MarkFlagRequired("backup-file")
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a Semaphore project from backup",
	Run: func(cmd *cobra.Command, args []string) {
		httpClient := createHttpClient(cmd)
		s := semaphore.New(httpClient)
		projectID := cmd.Flag("project").Value.String()
		s.Restore(projectID)
	},
}
