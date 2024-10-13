package cmd

import (
	"github.com/socheatsok78/semaphore-cli/semaphore"
	"github.com/spf13/cobra"
)

var (
	configBackupProjectID string
	configBackupFile      string
)

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVar(&configBackupProjectID, "project-id", "", "semaphore project id")
	backupCmd.Flags().StringVar(&configBackupFile, "output", "backup-%s.json", "output file")
	backupCmd.MarkFlagRequired("project")
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup of Semaphore project",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := semaphore.New(configAddr, configDNS)
		if err != nil {
			return err
		}
		err = s.Authenticate(configUsername, configPassword)
		if err != nil {
			return err
		}
		if err := s.Backup(configBackupProjectID, configBackupFile); err != nil {
			return err
		}
		return nil
	},
}
