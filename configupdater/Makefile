export GOPROXY=https://proxy.golang.org

GO ?= go
GOBIN := $(shell $(GO) env GOBIN)
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
FIRST_GOPATH := $(firstword $(subst :, ,$(GOPATH)))
GOPKGDIR := $(FIRST_GOPATH)/src/$(PROJECT)
GOPKGBASEDIR ?= $(shell dirname "$(GOPKGDIR)")

ifeq ($(GOBIN),)
GOBIN := $(FIRST_GOPATH)/bin
endif

GO_BUILD=$(GO) build

###
### Primary entry-point targets
###

default: test binaries

vendor:
	GO111MODULE=on $(GO) mod tidy
	GO111MODULE=on $(GO) mod vendor
	GO111MODULE=on $(GO) mod verify

gofmt:
	find . -name '*.go' ! -path './vendor/*' -exec gofmt -s -w {} \+

binaries: clean
	CGO_ENABLED=0 $(GO_BUILD) -o bin/configupdater

clean:
	rm -fr bin/

test:
	$(GO) test -v -count=0

build-image:
	docker build -t jeyaganesh/configupdater:0.0.1 .