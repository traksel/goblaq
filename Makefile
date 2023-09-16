INSTALL_PATH ?= /usr/local/bin
BINDIR 		 ?= $(CURDIR)/bin
BINNAME 	 ?= goblaq
SRC 		 := $(shell find . -type f -name '*.go' -print) go.mod go.sum
CGO_ENABLED  ?= 0
GOFLAGS 	 :=
SHELL        = /usr/bin/env bash
SERVICE_PATH ?= /lib/systemd/system/$(BINNAME).service

define SERVICE_BODY
[Unit]
Description=https://github.com/traksel/goblaq
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/$(BINNAME) daemon daemon
Restart=always

[Install]
WantedBy=multi-user.target
endef
export SERVICE_BODY

.PHONY: all
all: build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on CGO_ENABLED=$(CGO_ENABLED) go build -o '$(BINDIR)'/$(BINNAME) ./cmd/goblaq

.PHONY: install
install: build
	@mkdir -p $(HOME)/.goblaq
	@install "$(BINDIR)/$(BINNAME)" "$(INSTALL_PATH)/$(BINNAME)"

.PHONY: daemon
daemon:
	@echo "$$SERVICE_BODY" > "$(SERVICE_PATH)"


.PHONY: clean
clean:
	@rm -rf '$(BINDIR)'
