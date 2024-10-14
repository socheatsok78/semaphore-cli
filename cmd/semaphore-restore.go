package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-kit/log/level"
	"github.com/socheatsok78/semaphore-cli/internals"
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
		level.Info(internals.Logger).Log("msg", "Reading backup from file", "file", configRestoreBackupFile)
		file, err := os.Open(configRestoreBackupFile)
		if err != nil {
			return err
		}
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		backup := &semaphore.SemaphoreBackup{}
		if err = json.Unmarshal(fileBytes, backup); err != nil {
			return err
		}

		level.Info(internals.Logger).Log("msg", "Connecting to Semaphore")
		s, err := semaphore.New(configAddr, configDNS)
		if err != nil {
			return err
		}

		level.Info(internals.Logger).Log("msg", "Authenticating")
		err = s.Authenticate(configUsername, configPassword)
		if err != nil {
			return err
		}

		level.Info(internals.Logger).Log("msg", "Restoring project", "project", configRestoreProjectID)
		if err := s.Restore(configRestoreProjectID, backup); err != nil {
			return err
		}
		return nil
	},
}
