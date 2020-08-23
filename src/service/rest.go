package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 30,
	}
}

// Delete is the default implementation of a DELETE Request based on a default client and
// given valid uri.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func Delete(uri string, headers map[string]string, chbs chan<- []byte, cher chan<- error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", constants.BaseURI, uri), nil)
	if err != nil {
		cher <- err
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		cher <- err
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cher <- err
		return
	}

	if resp.StatusCode/100 != 2 {
		er := new(model.ErrorResponse)
		message := string(bs)
		if err := json.Unmarshal(bs, er); err != nil {
			message = er.Message
		}
		cher <- fmt.Errorf(constants.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
		return
	}

	chbs <- bs
}

// Post is the default implementation of a POST Request based on a default client and
// given valid uri, headers and body.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func Post(uri string, headers map[string]string, body []byte, chbs chan<- []byte, cher chan<- error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", constants.BaseURI, uri), bytes.NewBuffer(body))
	if err != nil {
		cher <- err
		return
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		cher <- err
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cher <- err
		return
	}

	if resp.StatusCode/100 != 2 {
		er := new(model.ErrorResponse)
		message := string(bs)
		if err := json.Unmarshal(bs, er); err == nil {
			message = er.Message
		}
		cher <- fmt.Errorf(constants.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
		return
	}

	chbs <- bs
}

// Get is the default implementation of a GET Request based on a default client and
// given valid uri and headers.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func Get(uri string, headers map[string]string, chbs chan<- []byte, cher chan<- error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", constants.BaseURI, uri), nil)
	if err != nil {
		cher <- err
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		cher <- err
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cher <- err
		return
	}

	if resp.StatusCode/100 != 2 {
		er := new(model.ErrorResponse)
		message := string(bs)
		if err := json.Unmarshal(bs, er); err == nil {
			message = er.Message
		}
		cher <- fmt.Errorf(constants.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
		return
	}

	chbs <- bs
}
