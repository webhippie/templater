DIST := dist
IMPORT := github.com/webhippie/templater

ifeq ($(OS), Windows_NT)
	EXECUTABLE := templater.exe
else
	EXECUTABLE := templater
endif

SHA := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u '+%Y%m%d')
LDFLAGS += -s -w -extldflags "-static" -X "$(IMPORT)/config.VersionDev=$(SHA)" -X "$(IMPORT)/config.VersionDate=$(DATE)"

TARGETS ?= linux/*,darwin/*,windows/*
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*")

TAGS ?=

ifneq ($(DRONE_TAG),)
	VERSION ?= $(subst v,,$(DRONE_TAG))
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(subst release/v,,$(DRONE_BRANCH))
	else
		VERSION ?= master
	endif
endif

.PHONY: all
all: build

.PHONY: update
update:
	@which govend > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/govend/govend; \
	fi
	govend -vtlu --prune

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(EXECUTABLE) $(DIST)

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: misspell
misspell:
	@which misspell > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell $(SOURCES)

.PHONY: generate
generate:
	go generate $(PACKAGES)

.PHONY: staticcheck
staticcheck:
	@which staticcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get honnef.co/go/staticcheck/cmd/staticcheck; \
	fi
	staticcheck $(PACKAGES)

.PHONY: errcheck
errcheck:
	@which errcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

.PHONY: varcheck
varcheck:
	@which varcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/opennota/check/cmd/varcheck; \
	fi
	varcheck $(PACKAGES)

.PHONY: structcheck
structcheck:
	@which structcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/opennota/check/cmd/structcheck; \
	fi
	structcheck $(PACKAGES)

.PHONY: unconvert
unconvert:
	@which unconvert > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mdempsky/unconvert; \
	fi
	unconvert $(PACKAGES)

.PHONY: interfacer
interfacer:
	@which interfacer > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mvdan/interfacer/cmd/interfacer; \
	fi
	interfacer $(PACKAGES)

.PHONY: ineffassign
ineffassign:
	@which ineffassign > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/gordonklaus/ineffassign; \
	fi
	ineffassign .

.PHONY: dupl
dupl:
	@which dupl > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mibk/dupl; \
	fi
	dupl .

.PHONY: lint
lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: test
test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

.PHONY: check
check: test

.PHONY: install
install: $(SOURCES)
	go install -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)'

.PHONY: build
build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	go build -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@

.PHONY: release
release: release-dirs release-build release-copy release-check

.PHONY: release-dirs
release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

.PHONY: release-build
release-build:
	@which xgo > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -targets '$(TARGETS)' -out $(EXECUTABLE)-$(VERSION) $(IMPORT)
ifeq ($(CI),drone)
	mv /build/* $(DIST)/binaries
endif

.PHONY: release-copy
release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

.PHONY: release-check
release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

.PHONY: updater
updater: updater-dirs updater-build updater-copy updater-push

.PHONY: updater-dirs
updater-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/updater

.PHONY: updater-build
updater-build:
	@which xgo > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -targets '$(TARGETS)' -out $(EXECUTABLE)-$(VERSION) $(IMPORT)
ifeq ($(CI),drone)
	mv /build/* $(DIST)/binaries
endif

.PHONY: updater-copy
updater-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/updater/$(word 3,$(subst -, ,$(notdir $(file))))-$(subst .exe,,$(word 4,$(subst -, ,$(notdir $(file)))));)

.PHONY: updater-push
updater-push:
	@which go-selfupdate > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/sanbornm/go-selfupdate; \
	fi
	go-selfupdate -o $(DIST)/publish $(DIST)/updater $(VERSION)

.PHONY: publish
publish: release updater
