
PROJECT=hblend
PWD=$(shell pwd)
GOPATH=$(PWD)
GOBIN=$(PWD)/bin
GOPKG=$(PWD)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := build_one

.PHONY: all build build_one clean dependencies test coverage

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

test:
	$(GOCMD) test ./src/$(PROJECT)/... -cover

coverage:
	rm -fr coverage
	mkdir -p coverage
	$(GOCMD) list $(PROJECT)/... > coverage/packages
	@i=a ; \
	while read -r P; do \
		i=a$$i ; \
		$(GOCMD) test ./src/$$P -cover -covermode=count -coverprofile=coverage/$$i.out; \
	done <coverage/packages

	echo "mode: count" > coverage/coverage
	cat coverage/*.out | grep -v "mode: count" >> coverage/coverage
	$(GOCMD) tool cover -html=coverage/coverage

