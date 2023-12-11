.DEFAULT_GOAL := build
SOURCE := $(wildcard *.go **/*.go)
GO_PATH := $(shell go env GOPATH)
INSTALL_PATH := $(GO_PATH)/bin/advent-of-code-2023
LOG_LEVEL ?= warn

.PHONY:
session:
ifndef AOC_SESSION_TOKEN
	$(error AOC_SESSION_TOKEN is undefined)
endif

.PHONY:
test: session
	go test common/*.go
	@#go test day00/*.go
	go test day01/*.go

.PHONY:
build: advent-of-code-2023

advent-of-code-2023: $(SOURCE)
	go build

.PHONY:
clean: 
	go clean

.PHONY:
run: build session
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2023 $(DAY) $(INPUT_FILE)

.PHONY:
run-all: build session
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2023 0
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2023 1

.PHONY:
install: $(INSTALL_PATH)

$(INSTALL_PATH): $(SOURCE)
	go install 

.PHONY:
uninstall:
	rm -f $(INSTALL_PATH)

.PHONY:
setup: go.mod
	go mod tidy
	go mod download

.PHONY:
go.mod: /usr/local/go/bin/go ~/.go
	go mod init

/usr/local/go/bin/go:
	sudo wget -c https://go.dev/dl/go1.21.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

~/.go:
	echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee $(HOME)/.go
	echo -e "\n. \$$HOME/.go" | tee -a $(HOME)/.bashrc
	