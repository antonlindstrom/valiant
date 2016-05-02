PROGRAM:=valiant

all: $(PROGRAM)

$(PROGRAM):
	go build -o $(PROGRAM) ./cmd/...

.PHONY: install
install: $(PROGRAM)
	cp $(PROGRAM) $(GOPATH)/bin/$(PROGRAM)

.PHONY: clean
clean:
	rm $(PROGRAM)
