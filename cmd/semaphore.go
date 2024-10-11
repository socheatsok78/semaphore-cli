package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	// Version is the version of the CLI
	Version = "dev"

	// Command-line root command
	rootCmd = &cobra.Command{
		Use:     "semaphore",
		Short:   "A backup and restore tool for Semaphore CI",
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			username := cmd.Flag("username").Value.String()
			if username == "" {
				username, err := readUsername()
				if err != nil {
					return err
				}
				cmd.Flags().Set("username", username)
			}

			password := cmd.Flag("password").Value.String()
			if password == "" {
				password, err := readPassword()
				if err != nil {
					return err
				}
				cmd.Flags().Set("password", password)
			}

			return nil
		},
	}
)

func init() {
	rootCmd.PersistentFlags().String("dns", "127.0.0.11", "dns resolver")
	rootCmd.PersistentFlags().String("addr", "https://cloud.semaphoreui.com", "the semaphore address")
	rootCmd.PersistentFlags().String("username", "", "the semaphore username")
	rootCmd.PersistentFlags().String("password", "", "the semaphore password")
}

func readUsername() (string, error) {
	fmt.Print("Enter username: ")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		return "", fmt.Errorf("username cannot be empty")
	}
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	return username, nil
}

func readPassword() (string, error) {
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	if len(bytePassword) == 0 {
		return "", fmt.Errorf("password cannot be empty")
	}
	fmt.Println()
	return string(bytePassword), nil
}

func Execute() error {
	return rootCmd.Execute()
}
