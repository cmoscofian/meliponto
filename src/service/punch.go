package service

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/util/constants"
)

func GetPunchByID(token, punchID string, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s/%s", constants.SitePunchURI, punchID)

	Get(uri, headers, chbs, cher)
}

func GetPunchByDateRange(token string, start, end time.Time, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s?inicio=%s&fim=%s", constants.SitePunchURI, start.Format(constants.DateLayout), end.Format(constants.DateLayout))

	Get(uri, headers, chbs, cher)
}

func PostPunch(token string, body []byte, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}

	Post(constants.SitePunchURI, headers, body, chbs, cher)
}
