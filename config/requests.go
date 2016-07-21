package config

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// SendRequest sends a request to a server and returns the test response.
func (t *TestSpec) SendRequest(addr string) (*Response, error) {
	reqURL := addr + t.Request.Path

	body := strings.NewReader(t.Request.Body)
	req, err := http.NewRequest(t.Request.Method, reqURL, body)
	if err != nil {
		return nil, err
	}

	// Add headers to request.
	for k, v := range t.Request.Headers {
		req.Header.Add(k, v)
	}

	client := http.DefaultClient
	client.Timeout = time.Second * time.Duration(t.TimeoutSeconds)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Make sure we handle the error of resp.Body.Close().
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Add headers, we're not allowing multiple values for a header.
	headers := make(map[string]string)
	for k, vals := range resp.Header {
		if len(vals) > 0 {
			headers[k] = vals[0]
		}
	}

	return &Response{
		Code:    resp.StatusCode,
		Body:    string(respBody),
		Headers: headers,
	}, err
}
