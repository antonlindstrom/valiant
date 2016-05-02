package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
	"gopkg.in/yaml.v2"
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

	defer resp.Body.Close()
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
	}, nil
}

// Update returns the generated config from the TestSpec
func (t *TestSpec) Update() ([]byte, error) {

	// Remove headers that changes between requests.
	for k, _ := range t.Response.Headers {
		switch strings.ToLower(k) {
		case "date":
			delete(t.Response.Headers, k)
		case "last-modified":
			delete(t.Response.Headers, k)
		case "expires":
			delete(t.Response.Headers, k)
		}
	}

	return yaml.Marshal(t)
}

// Validate runs trough the expected response against the actual response and
// returns an array with issues.
func (t *TestSpec) Validate(response Response) ([]string, error) {
	var issues []string

	if response.Code != t.Response.Code {
		s := fmt.Sprintf("response code mismatch: got %d, want %d", response.Code, t.Response.Code)
		issues = append(issues, s)
	}

	for k, v := range t.Response.Headers {
		value, ok := response.Headers[k]
		if !ok {
			issues = append(issues, fmt.Sprintf("response header mismatch: %q not found", k))
		}

		if value != v {
			issues = append(issues, fmt.Sprintf("response header mismatch: %q value want %q, got %q", k, v, value))
		}
	}

	if response.Body != t.Response.Body {
		diff := diffmatchpatch.New().DiffMain(t.Response.Body, response.Body, true)

		var resp []string
		resp = append(resp, "response body mismatch:\n")

		for _, d := range diff {
			switch d.Type {
			case diffmatchpatch.DiffDelete:
				resp = append(resp, "\x1b[31m"+d.Text+"\x1b[0m")
			case diffmatchpatch.DiffInsert:
				resp = append(resp, "\x1b[32m"+d.Text+"\x1b[0m")
			case diffmatchpatch.DiffEqual:
				resp = append(resp, "\x1b[0m"+d.Text+"\x1b[0m")
			}
		}

		issues = append(issues, resp...)
	}

	return issues, nil
}

// Parse []byte config to TestSpec.
func Parse(b []byte) (*TestSpec, error) {
	var t TestSpec
	err := yaml.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}

	if t.TimeoutSeconds == 0 {
		t.TimeoutSeconds = 3
	}

	return &t, nil
}

// GenerateExample generates example config YAML output.
func GenerateExample() ([]byte, error) {
	test := &TestSpec{
		TimeoutSeconds: 3,
		Request: Request{
			Method: "GET",
			Path:   "/",
			Headers: map[string]string{
				"Host": "example.com",
			},
		},
		Response: Response{
			Headers: map[string]string{
				"Content-Type": "encoding/json",
			},
			Body: "Example body",
			Code: 200,
		},
	}

	return yaml.Marshal(test)
}
