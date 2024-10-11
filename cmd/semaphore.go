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

	// Flag variables
	usernameStdinFlag bool
	passwordStdinFlag bool

	// Command-line root command
	rootCmd = &cobra.Command{
		Use:     "semaphore",
		Short:   "A backup and restore tool for Semaphore CI",
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if usernameStdinFlag {
				username, err := readUsername()
				if err != nil {
					return err
				}
				cmd.Flags().Set("username", username)
			}
			if passwordStdinFlag {
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

	// Username flags
	rootCmd.PersistentFlags().String("username", "", "the semaphore username")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.PersistentFlags().BoolVar(&usernameStdinFlag, "username-stdin", false, "read username from stdin")
	rootCmd.MarkFlagsOneRequired("username", "username-stdin")

	// Password flags
	rootCmd.PersistentFlags().String("password", "", "the semaphore password")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.PersistentFlags().BoolVar(&passwordStdinFlag, "password-stdin", false, "read password from stdin")
	rootCmd.MarkFlagsOneRequired("password", "password-stdin")
}

func readUsername() (string, error) {
	fmt.Print("Enter username: ")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		return "", err
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
