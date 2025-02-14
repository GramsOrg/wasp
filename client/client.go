package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/iotaledger/wasp/packages/webapi/v1/model"
)

var ErrNotAuthorized = errors.New("unauthorized request rejected")

// WaspClient allows to make requests to the Wasp web API.
type WaspClient struct {
	httpClient http.Client
	baseURL    string
	token      string
	logFunc    func(msg string, args ...interface{})
}

// NewWaspClient returns a new *WaspClient with the given baseURL and httpClient.
func NewWaspClient(baseURL string, httpClient ...http.Client) *WaspClient {
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	if len(httpClient) > 0 {
		return &WaspClient{baseURL: baseURL, httpClient: httpClient[0]}
	}
	return &WaspClient{baseURL: baseURL}
}

func (c *WaspClient) WithLogFunc(logFunc func(msg string, args ...interface{})) *WaspClient {
	c.logFunc = logFunc
	return c
}

func (c *WaspClient) WithToken(token string) *WaspClient {
	if len(token) > 0 {
		c.token = token
	}

	return c
}

func (c *WaspClient) log(msg string, args ...interface{}) {
	if c.logFunc == nil {
		return
	}
	c.logFunc(msg, args...)
}

func processResponse(res *http.Response, decodeTo interface{}) error {
	if res == nil || res.Body == nil {
		return errors.New("unable to read response body")
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return ErrNotAuthorized
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body: %w", err)
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if decodeTo != nil {
			return json.Unmarshal(resBody, decodeTo)
		}
		return nil
	}

	errRes := &model.HTTPError{}
	if err := json.Unmarshal(resBody, errRes); err != nil {
		errRes.Message = http.StatusText(res.StatusCode)
	}
	errRes.StatusCode = res.StatusCode
	errRes.Message = string(resBody)
	return errRes
}

func (c *WaspClient) do(method, route string, reqObj, resObj interface{}) error {
	// marshal request object
	var data []byte
	if reqObj != nil {
		var err error
		data, err = json.Marshal(reqObj)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
	}

	// construct request
	url := fmt.Sprintf("%s/%s", strings.TrimRight(c.baseURL, "/"), strings.TrimLeft(route, "/"))
	req, err := http.NewRequestWithContext(context.Background(), method, url, func() io.Reader {
		if data == nil {
			return nil
		}
		return bytes.NewReader(data)
	}())
	if err != nil {
		return fmt.Errorf("http.NewRequest [%s %s]: %w", method, url, err)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.token))
	}

	// make the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %s: %w", method, url, err)
	}

	// write response into response object
	err = processResponse(res, resObj)
	if err != nil {
		return fmt.Errorf("%s %s: %w", method, url, err)
	}
	return nil
}

// BaseURL returns the baseURL of the client.
func (c *WaspClient) BaseURL() string {
	return c.baseURL
}
