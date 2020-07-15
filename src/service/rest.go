package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cmoscofian/meliponto/src/util/constants"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 30,
	}
}

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

	chbs <- bs
}

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

	chbs <- bs
}
