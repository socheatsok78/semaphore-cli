package internals

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func ReadUsername() (string, error) {
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

func ReadPassword() (string, error) {
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
