package config

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	readFile := func(f string) []byte {
		file, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatalf("failed to read file %q", f)
		}

		return file
	}

	tests := []struct {
		file []byte
		spec *TestSpec
	}{
		{
			file: readFile("01_simple_test.yml"),
			spec: &TestSpec{
				Request: Request{
					Method: "GET",
					Path:   "/",
					Headers: map[string]string{
						"Host": "example.com",
					},
				},
				Response: Response{
					Code: 200,
					Body: "",
				},
				TimeoutSeconds:   3,
				SkipValidateBody: false,
			},
		},
		{
			file: readFile("02_simple_test.yml"),
			spec: &TestSpec{
				Request: Request{
					Method: "GET",
					Path:   "/",
					Headers: map[string]string{
						"Host": "example.com",
					},
				},
				Response: Response{
					Code: 200,
					Body: "",
				},
				TimeoutSeconds:   3,
				SkipValidateBody: true,
			},
		},
	}

	for _, test := range tests {
		spec, err := Parse(test.file)
		if err != nil {
			t.Fatalf("want err = nil, got err = %v", err)
		}

		if !reflect.DeepEqual(*test.spec, *spec) {
			t.Errorf("want spec = %+v, got spec = %+v", *test.spec, *spec)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		spec *TestSpec
		res  *TestSpec
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
				Response: Response{
					Code: 204,
					Headers: map[string]string{
						"date":          "123",
						"last-modified": "123",
						"expires":       "123",
						"x-testing":     "hello",
					},
					Body: "testing",
				},
				TimeoutSeconds:   3,
				SkipValidateBody: false,
			},
			res: &TestSpec{
				Request: Request{
					Method: "GET",
					Path:   "/",
					Headers: map[string]string{
						"Host": "example.com",
					},
				},
				Response: Response{
					Code: 204,
					Headers: map[string]string{
						"x-testing": "hello",
					},
					Body: "testing",
				},
				TimeoutSeconds:   3,
				SkipValidateBody: false,
			},
		},
	}

	for _, test := range tests {
		b, err := test.spec.Update()
		if err != nil {
			t.Fatalf("want err = nil, got err = %v", err)
		}

		spec, err := Parse(b)
		if err != nil {
			t.Fatalf("want err = nil, got err = %v", err)
		}

		if !reflect.DeepEqual(test.res, spec) {
			t.Errorf("want spec = %+v, got spec = %+v", *test.res, *spec)
		}
	}
}
