
PROJECT=hblend
GOPATH=$(shell pwd)/_vendor
GOBIN=$(GOPATH)/bin
GOPKG=$(GOPATH)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := build_one

.PHONY: all build clean dependencies setup

all: build

clean:
	rm -fr _vendor

setup:
	mkdir -p _vendor/src
	ln -s ../.. _vendor/src/hblend

dependencies:
	$(GOCMD) get $(PROJECT)

build_one: clean setup dependencies
	$(GOCMD) build -o $(GOBIN)/$(PROJECT);

build: clean setup dependencies
	for GOOS in "windows" "linux" "darwin"; do \
		for GOARCH in "386" "amd64"; do \
			echo "Building $$GOOS-$$GOARCH..."; \
			echo "GOOS=$$GOOS GOARCH=$$GOARCH $(GOCMD) build -o $(GOBIN)/$(PROJECT).$$GOOS.$$GOARCH"; \
			GOOS=$$GOOS GOARCH=$$GOARCH $(GOCMD) build -o $(GOBIN)/$(PROJECT).$$GOOS.$$GOARCH; \
		done \
	done
	ls $(GOBIN)
