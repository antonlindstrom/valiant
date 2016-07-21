package config

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

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
