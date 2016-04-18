DIST := dist
BIN := bin
EXECUTABLE := templater
SHA := $(shell git rev-parse --short HEAD)

RELEASES ?= $(BIN)/$(EXECUTABLE)-linux-amd64 \
	$(BIN)/$(EXECUTABLE)-linux-386 \
	$(BIN)/$(EXECUTABLE)-linux-arm \
	$(BIN)/$(EXECUTABLE)-linux-arm64 \
	$(BIN)/$(EXECUTABLE)-darwin-amd64

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(DRONE_BRANCH)
	else
		VERSION ?= master
	endif
endif

LDFLAGS += -X "github.com/webhippie/templater/config.VersionDev=$(SHA)"

all: clean deps build test

clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/govend/govend
	govend -v

vendor:
	govend -vtlu

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

lint:
	for PKG in $(PACKAGES); do golint $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

build: $(BIN)/$(EXECUTABLE)

release: $(RELEASES)

updater:
	go get -u github.com/sanbornm/go-selfupdate
	go-selfupdate -o $(DIST)/publish $(DIST)/updater $(VERSION)

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

$(BIN)/$(EXECUTABLE): $(wildcard *.go)
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(BIN)/$(EXECUTABLE)-%: GOOS=$(word 1,$(subst -, ,$*))
$(BIN)/$(EXECUTABLE)-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(BIN)/$(EXECUTABLE)-%:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags '-s -w $(LDFLAGS)' -o $@
	mkdir -p $(DIST)/updater
	cp $@ $(DIST)/updater/$(GOOS)-$(GOARCH)
	mkdir -p $(DIST)/release
	cp $@ $(DIST)/release/$(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH)
	cd $(DIST)/release && sha256sum $(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH) > $(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH).sha256
	mkdir -p $(DIST)/latest
	cp $@ $(DIST)/latest/$(EXECUTABLE)-latest-$(GOOS)-$(GOARCH)
	cd $(DIST)/latest && sha256sum $(EXECUTABLE)-latest-$(GOOS)-$(GOARCH) > $(EXECUTABLE)-latest-$(GOOS)-$(GOARCH).sha256

.PHONY: all clean deps vendor fmt vet lint test build updater install
.PRECIOUS: $(BIN)/$(EXECUTABLE)-%
