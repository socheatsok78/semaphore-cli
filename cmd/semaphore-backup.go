package cmd

import (
	"io"
	"os"

	"github.com/go-kit/log/level"
	"github.com/socheatsok78/semaphore-cli/internals"
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
	backupCmd.Flags().StringVar(&configBackupFile, "output", "stdout", "output")
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
		level.Info(internals.Logger).Log("msg", "Creating backup", "project", configBackupProjectID)
		data, err := s.Backup(configBackupProjectID, configBackupFile)
		if err != nil {
			return err
		}
		defer data.Close()
		if configBackupFile == "stdout" {
			level.Info(internals.Logger).Log("msg", "Writing backup to stdout")
			io.Copy(os.Stdout, data)
		} else {
			file, err := os.Create(configBackupFile)
			if err != nil {
				return err
			}
			defer file.Close()
			level.Info(internals.Logger).Log("msg", "Writing backup to file", "file", configBackupFile)
			io.Copy(file, data)
		}
		return nil
	},
}
