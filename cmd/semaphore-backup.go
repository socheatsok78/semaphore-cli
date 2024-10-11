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
		addr := cmd.Flag("addr").Value.String()
		dns := cmd.Flag("dns").Value.String()
		s := semaphore.New(addr, dns)
		err := s.Login(cmd.Flag("username").Value.String(), cmd.Flag("password").Value.String())
		if err != nil {
			panic(err)
		}
		projectID := cmd.Flag("project").Value.String()
		err = s.Backup(projectID)
		if err != nil {
			panic(err)
		}
	},
}
