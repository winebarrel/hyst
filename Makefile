VERSION = $(shell git tag | tail -n 1)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
RUNTIME_GOPATH := $(GOPATH):$(shell pwd)
SRC := $(wildcard *.go) $(wildcard src/hyst/*.go)
TEST_SRC := $(wildcard src/hyst/*_test.go)

all: hyst

hyst: $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build

test: $(TEST_SRC)
	GOPATH=$(RUNTIME_GOPATH) go test -v $(TEST_SRC)

dev_dependency:
	go get github.com/stretchr/testify

clean:
	rm -f hyst{,.exe} *.gz *.zip

package: clean hyst
ifeq ($(GOOS),windows)
	zip hyst-$(VERSION)-$(GOOS)-$(GOARCH).zip hyst.exe
else
	gzip -c hyst > hyst-$(VERSION)-$(GOOS)-$(GOARCH).gz
endif
