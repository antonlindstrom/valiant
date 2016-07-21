package config

type Request struct {
	Headers map[string]string `yaml:"headers"`
	Method  string            `yaml:"method,omitempty"`
	Path    string            `yaml:"path,omitempty"`
	Body    string            `yaml:"body,omitempty"`
}

type Response struct {
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body,omitempty"`
	Code    int               `yaml:"code,omitempty"`
}

type TestSpec struct {
	Request          Request  `yaml:"request"`
	Response         Response `yaml:"response"`
	TimeoutSeconds   int      `yaml:"timeout_seconds"`
	SkipValidateBody bool     `yaml:"skip_validate_body"`
}
