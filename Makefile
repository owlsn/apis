-include .env

# VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
ifndef $(VERSION)
	VERSION := $(shell git describe --always)
endif
# Go related variables.
BASE := $(shell pwd)
GOBASE := $(BASE)/vendor:$(BASE)
GOBIN := $(BASE)/bin
GOFILES := $(wildcard *.go)

# Use linker flags to provide version/build settings
# LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# 查看变量是否设置正确
variables:
	@echo " VERSION:$(VERSION) "
	@echo " BUILD:$(BUILD) "
	@echo " PROJECTNAME:$(PROJECTNAME) "
	@echo " BASE:$(BASE)"
	@echo " PID:$(PID) "

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## format all code
fmt: go-fmt

## start: Start in development mode. Auto-starts when code changes.
start:
	@bash -c "$(MAKE) clean compile start-server run='make clean compile start-server'"
  # @bash -c "trap 'make stop' EXIT; $(MAKE) clean compile start-server run='make clean compile start-server'"

## stop: Stop development mode.
stop: stop-server

start-server:
ifneq ($(PID), $(wildcard $(PID)))
	@$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"
endif
	@echo "  >  $(PROJECTNAME) is available at `cat $(PID)`"

stop-server:
ifeq ($(PID), $(wildcard $(PID)))
	@kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)
endif

## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
# watch:
# 	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) yolo -i . -e vendor -e bin -c "$(run)"

restart-server: stop-server start-server

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## exec: Run given command, wrapped with custom GOBASE. e.g; make exec run="go test ./..."
exec:
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-$(MAKE) go-clean

go-compile: go-get go-build

go-build:
	@echo "  >  Building binary..."
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	# @GOBASE=$(GOBASE) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-generate:
	@echo "  >  Generating dependency files..."
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) go get $(get)

go-install:
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOBASE=$(GOBASE) GOBIN=$(GOBIN) go clean

go-fmt:
	@go fmt $(BASE)/...

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo