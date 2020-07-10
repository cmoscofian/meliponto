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
var timeout time.Duration

func init() {
	client = &http.Client{
		Timeout: timeout,
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

func Get(uri string, headers map[string]string, ch chan<- []byte) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", constants.BaseURI, uri), nil)
	if err != nil {
		ch <- []byte(err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		ch <- []byte(err.Error())
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []byte(err.Error())
		return
	}

	ch <- bs
}
