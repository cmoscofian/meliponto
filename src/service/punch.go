package service

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/util/constants"
)

func Punch(token string, body []byte, chbs chan<- []byte, cher chan<- error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}

	Post(constants.SitePunchURI, headers, body, chbs, cher)
}
