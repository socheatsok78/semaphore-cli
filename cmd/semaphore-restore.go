package cmd

import (
	"github.com/socheatsok78/semaphore-cli/semaphore"
	"github.com/spf13/cobra"
)

var (
	configRestoreProjectID  string
	configRestoreBackupFile string
)

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVar(&configRestoreProjectID, "project-id", "", "semaphore project id")
	restoreCmd.Flags().StringVar(&configRestoreBackupFile, "backup-file", "", "semaphore backup file to restore from")
	restoreCmd.MarkFlagRequired("backup-file")
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a Semaphore project from backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := semaphore.New(configAddr, configDNS)
		if err != nil {
			return err
		}
		err = s.Authenticate(configUsername, configPassword)
		if err != nil {
			return err
		}
		if err := s.Restore(configRestoreProjectID, configRestoreBackupFile); err != nil {
			return err
		}
		return nil
	},
}
