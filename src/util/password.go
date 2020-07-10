package util

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword reads a password from stdin
func GetPassword() ([]byte, error) {
	var password []byte
	var err error

	counter := 0

	for len(password) < 1 {
		if counter > 0 {
			fmt.Printf("Sorry, password cannot be empty\n\n")
		}

		fmt.Print("Enter your password: ")

		password, err = terminal.ReadPassword(0)
		if err != nil {
			return nil, err
		}

		fmt.Println()
		counter++
	}

	return password, nil
}
