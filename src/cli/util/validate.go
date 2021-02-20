package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// AllowPunch generates an interface to be dealt by the user
// whether or not it should continue the process on the face of
// previous values found on the same query period.
func AllowPunch(query *entity.QueryPunchResponse) {
	if query.HasData() {
		scanner := bufio.NewScanner(os.Stdin)
		counter := 0

		for {
			if counter > 0 {
				fmt.Printf("Invalid option, please insert a valid option!\n\n")
			}

			fmt.Printf("[\u001b[33mWARNING\u001b[0m] You already have %d punches in this period! Would you like to continue? (y/N) ", query.Total)
			scanner.Scan()
			input := strings.ToLower(scanner.Text())

			if input == "y" || input == "yes" {
				break
			}

			if input == "n" || input == "no" || input == "" {
				fmt.Println("Exiting... NOTHING has been punched!")
				os.Exit(1)
			}

			counter++
		}
	}
}
