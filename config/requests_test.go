package config

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSendRequest(t *testing.T) {
	tests := []struct {
		spec *TestSpec
		resp *Response
	}{
		{
			spec: &TestSpec{
				Request: Request{
					Method: "GET",
					Path:   "/",
					Headers: map[string]string{
						"Host": "example.com",
					},
				},
				TimeoutSeconds: 3,
			},
			resp: &Response{
				Code: http.StatusOK,
				Body: "testing",
				Headers: map[string]string{
					"X-Testing":      "testing",
					"Content-Length": "7",
					"Content-Type":   "text/plain; charset=utf-8",
				},
			},
		},
	}

	// Set up test handler.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Testing", "testing")
		_, err := w.Write([]byte("testing"))
		if err != nil {
			t.Fatalf("want err = nil, got err = %v", err)
		}
	})

	for _, test := range tests {
		server := httptest.NewServer(handler)

		resp, err := test.spec.SendRequest(server.URL)
		if err != nil {
			t.Fatalf("want err = nil, got err = %v", err)
		}

		delete(resp.Headers, "Date")

		if !reflect.DeepEqual(test.resp, resp) {
			t.Fatalf("want resp = %+v, got resp = %+v", *test.resp, *resp)
		}

		server.Close()
	}
}
