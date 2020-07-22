package service

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/util/constants"
)

// GetPunchByID is responsible for building the request for fetching a single punch by it's
// ID given a valid context and making the get request to the punch URI.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func GetPunchByID(token, punchID string, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s/%s", constants.SitePunchURI, punchID)

	Get(uri, headers, chbs, cher)
}

// GetPunchByDateRange is responsible for building the request for fetching a range of punches
// given a valid context, start and end date and making the get request to the punch URI.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
// PS: To get a single day request just use the same date for start and end.
func GetPunchByDateRange(token string, start, end time.Time, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s?inicio=%s&fim=%s", constants.SitePunchURI, start.Format(constants.DateLayout), end.Format(constants.DateLayout))

	Get(uri, headers, chbs, cher)
}

// PostPunch is responsible for building the request for the punch given a valid context
// and body and making the post request to the login URI.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func PostPunch(token string, body []byte, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}

	Post(constants.SitePunchURI, headers, body, chbs, cher)
}
