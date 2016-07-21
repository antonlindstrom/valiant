PROGRAM:=valiant

GOFILES := $(shell find . -type f -name '*.go')

all: $(PROGRAM)

$(PROGRAM): $(GOFILES)
	go build -o $(PROGRAM) ./cmd/...

.PHONY: check
check: $(PROGRAM)
	go test -cover ./...
	go vet ./...
	errcheck ./...
	gofmt -s -l -w -d ./cmd/valiant ./config
	misspell -error ./cmd/valiant ./config

.PHONY: install
install: $(PROGRAM)
	cp $(PROGRAM) $(GOPATH)/bin/$(PROGRAM)

.PHONY: clean
clean:
	rm $(PROGRAM)

.PHONY: install-test-utils
install-test-utils:
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/kisielk/errcheck
