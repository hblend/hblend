
PROJECT=hblend
PWD=$(shell pwd)
GOPATH=$(PWD)
GOBIN=$(PWD)/bin
GOPKG=$(PWD)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := build_one

.PHONY: all build clean dependencies

all: build

clean:
	@echo "Cleaning..."
	rm -fr $(GOPKG)

dependencies:
	$(GOCMD) get $(PROJECT)

build_one: clean
	$(GOCMD) build -o $(GOBIN)/$(PROJECT) $(PROJECT)

build: clean
	@for GOOS in "windows" "linux" "darwin"; do \
		for GOARCH in "386" "amd64"; do \
			echo "Building $$GOOS-$$GOARCH..."; \
			GOOS=$$GOOS GOARCH=$$GOARCH $(GOCMD) build -o $(GOBIN)/$(PROJECT).$$GOOS.$$GOARCH $(PROJECT); \
		done \
	done
	ls $(GOBIN)
