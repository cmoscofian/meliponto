package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
)

// HandleFetchToPunch is responsible for interactively prompts
// the user to take an intentional action to proceed or to abort.
// It communicates with all other systems via channels
// ([]byte channel and error channel).
func HandleFetchToPunch(token string, start, end time.Time, chbs chan []byte, cher chan error) error {
	query, err := HandleFetch(token, start, end, chbs, cher)
	if err != nil {
		return err
	}

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
				return nil
			}

			if input == "n" || input == "no" || input == "" {
				fmt.Println("Exiting... NOTHING has been punched!")
				os.Exit(1)
			}

			counter++
		}
	}

	return nil
}

// HandleFetch is responsible for fetching all previous punches
// already registered within the range provided from the system.
// It communicates with all other systems via channels
// ([]byte channel and error channel).
func HandleFetch(token string, start, end time.Time, chbs chan []byte, cher chan error) (*model.Query, error) {
	query := new(model.Query)
	go service.GetPunchByDateRange(token, start, end, chbs, cher)

	select {
	case response := <-chbs:
		if err := json.Unmarshal(response, query); err != nil {
			return nil, err
		}
		return query, nil
	case err := <-cher:
		return nil, err
	}
}
