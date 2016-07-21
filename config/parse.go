package config

import (
	"strings"

	"gopkg.in/yaml.v2"
)

// Parse []byte config to TestSpec.
func Parse(b []byte) (*TestSpec, error) {
	var t TestSpec
	err := yaml.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}

	// Default to 3 seconds timeout.
	if t.TimeoutSeconds == 0 {
		t.TimeoutSeconds = 3
	}

	return &t, nil
}

// Update returns the generated config from the TestSpec.
func (t *TestSpec) Update() ([]byte, error) {
	// Remove headers that changes between requests.
	for k := range t.Response.Headers {
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
