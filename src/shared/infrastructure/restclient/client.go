package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// The Client type is an interface that defines
// the default implementation for a rest communication layer
// to be used across different clients.
type Client interface {
	Get(uri string, headers map[string]string, object interface{}) error
	Post(uri string, headers map[string]string, body, object interface{}) error
	Put(uri string, headers map[string]string, body, object interface{}) error
	Delete(uri string, headers map[string]string, object interface{}) error
}

type restClientPool struct {
	baseURI string
	client  *http.Client
	headers map[string]string
}

// NewRestClientPool returns a implementation of the RestClient
// interface to be used as a rest communication layer.
func NewRestClientPool(uri string, headers map[string]string, t time.Duration) Client {
	return &restClientPool{
		baseURI: uri,
		client: &http.Client{
			Timeout: t,
		},
		headers: headers,
	}
}

// Get is the default implementation of a GET Request based on a restClient and
// given valid uri, headers and object.
// The "object" param must be a pointer of an entity that matches the expected response.
// Returns an error
func (r *restClientPool) Get(uri string, headers map[string]string, object interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", r.baseURI, uri), nil)
	if err != nil {
		return err
	}

	// Set pool client headers
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// Set custom call headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		er := new(entity.ErrorResponse)
		message := string(bs)
		if err := json.Unmarshal(bs, er); err == nil {
			message = er.Message
		}
		return fmt.Errorf(constant.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
	}

	if err := json.Unmarshal(bs, object); err != nil {
		return err
	}

	return nil
}

// Post is the default implementation of a POST Request based on a restClient and
// given valid uri, headers, body and object.
// The "object" param must be a pointer of an entity that matches the expected response.
// Returns an error
func (r *restClientPool) Post(uri string, headers map[string]string, body, object interface{}) error {
	bbs, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", r.baseURI, uri), bytes.NewBuffer(bbs))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// Set pool client headers
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// Set custom call headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rbs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		er := new(entity.ErrorResponse)
		message := string(rbs)
		if err := json.Unmarshal(rbs, er); err == nil {
			message = er.Message
		}
		return fmt.Errorf(constant.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
	}

	if err := json.Unmarshal(rbs, object); err != nil {
		return err
	}

	return nil
}

// Put is the default implementation of a PUT Request based on a restClient and
// given valid uri, headers, body and object.
// The "object" param must be a pointer of an entity that matches the expected response.
// Returns an error
func (r *restClientPool) Put(uri string, headers map[string]string, body, object interface{}) error {
	bbs, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", r.baseURI, uri), bytes.NewBuffer(bbs))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// Set pool client headers
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// Set custom call headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rbs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		er := new(entity.ErrorResponse)
		message := string(rbs)
		if err := json.Unmarshal(rbs, er); err == nil {
			message = er.Message
		}
		return fmt.Errorf(constant.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
	}

	if err := json.Unmarshal(rbs, object); err != nil {
		return err
	}

	return nil
}

// Delete is the default implementation of a DELETE Request based on a restClient and
// given valid uri, headers and object.
// The "object" param must be a pointer of an entity that matches the expected response.
// Returns an error
func (r *restClientPool) Delete(uri string, headers map[string]string, object interface{}) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", r.baseURI, uri), nil)
	if err != nil {
		return err
	}

	// Set pool client headers
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// Set custom call headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		er := new(entity.ErrorResponse)
		message := string(bs)
		if err := json.Unmarshal(bs, er); err != nil {
			message = er.Message
		}
		return fmt.Errorf(constant.RestServiceError, http.StatusText(resp.StatusCode), resp.StatusCode, message)
	}

	if err := json.Unmarshal(bs, object); err != nil {
		return err
	}

	return nil
}
