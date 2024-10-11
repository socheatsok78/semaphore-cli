package cmd

import (
	"github.com/socheatsok78/semaphore-cli/semaphore"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().String("project-id", "", "semaphore project id")
	restoreCmd.Flags().String("backup-file", "", "semaphore backup file to restore from")
	restoreCmd.MarkFlagRequired("backup-file")
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a Semaphore project from backup",
	Run: func(cmd *cobra.Command, args []string) {
		addr := cmd.Flag("addr").Value.String()
		dns := cmd.Flag("dns").Value.String()
		s := semaphore.New(addr, dns)
		err := s.Login(cmd.Flag("username").Value.String(), cmd.Flag("password").Value.String())
		if err != nil {
			panic(err)
		}
		projectID := cmd.Flag("project-id").Value.String()
		s.Restore(projectID)
	},
}
