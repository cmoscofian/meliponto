package handlers

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
)

// HandleDelete is responsible for deleting all previous punches
// already registered within the range provided from the system.
// It communicates with all other systems via channels
// ([]byte channel and error channel).
func HandleDelete(token string, punches []*model.PunchResponse, chbs chan []byte, cher chan error) error {
	for _, p := range punches {
		go service.DeletePunchByID(token, p.ID, chbs, cher)
	}

	for range punches {
		select {
		case <-chbs:
			continue
		case err := <-cher:
			return err
		}
	}

	fmt.Println("Previous punches have been deleted!")

	return nil
}
