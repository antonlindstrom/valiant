PROGRAM:=valiant

GOFILES := $(shell find . -type f -name '*.go')

GOLANGCI_LINT_PREFIX=$(shell go env GOPATH)

all: $(PROGRAM)

$(PROGRAM): $(GOFILES)
	go build -o $(PROGRAM) ./cmd/...

.PHONY: check
check: $(PROGRAM)
	go test -v -cover ./...
	go vet ./...
	golangci-lint run ./...

.PHONY: install
install: $(PROGRAM)
	install -m 755 $(PROGRAM) $(GOPATH)/bin/$(PROGRAM)

.PHONY: clean
clean:
	rm $(PROGRAM)

.PHONY: install-test-utils
install-test-utils:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOLANGCI_LINT_PREFIX)/bin v1.24.0
