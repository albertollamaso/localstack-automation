-include .env

# set the shell to bash always
SHELL         := /bin/bash

# set make and shell flags to exit on errors
MAKEFLAGS     += --warn-undefined-variables
.SHELLFLAGS   := -euo pipefail -c

ARCH = amd64 arm64
BUILD_ARGS ?=

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)
GOCMD=go
GOTEST=$(GOCMD) test
VERSION?=0.0.0
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true
PROJECTNAME=localstack-automation

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt


# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# ====================================================================================
# Colors

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)
BLUE   := $(shell printf "\033[34m")
RED    := $(shell printf "\033[31m")
CNone  := $(shell printf "\033[0m")

# ====================================================================================
# Logger

TIME_LONG	= `date +%Y-%m-%d' '%H:%M:%S`
TIME_SHORT	= `date +%H:%M:%S`
TIME		= $(TIME_SHORT)

INFO	= echo ${TIME} ${BLUE}[ .. ]${CNone}
WARN	= echo ${TIME} ${YELLOW}[WARN]${CNone}
ERR		= echo ${TIME} ${RED}[FAIL]${CNone}
OK		= echo ${TIME} ${GREEN}[ OK ]${CNone}
FAIL	= (echo ${TIME} ${RED}[FAIL]${CNone} && false)

# ====================================================================================
# Actions

# create-secrets:

aws-list-secrets: ## List all AWS secrets from localstack
	./bin/localstack-automation --awssecretslist

# total-secrets:

# ====================================================================================
# Localstack
.PHONY: localstack-install
localstack-install: ## Install localstack
	python3 -m pip install localstack

.PHONY: localstack-start
localstack-start: ## Start localstack
	localstack start -d
	sleep 5

.PHONY: localstack-status
localstack-status: ## Check localstack status
	localstack status services

.PHONY: localstack-test-awscli
localstack-test-awscli: ## test AWS CLI locally with localstack using AWS Secrets
	rand=test-$$RANDOM && \
	echo "Creating AWS Secret: $$rand"  && \
	aws --endpoint-url=http://localhost:4566 secretsmanager create-secret --name $$rand --secret-string file://secretvalues.json
	sleep 5
	@echo "  >  Listing AWS secrets"
	aws --endpoint-url=http://localhost:4566 secretsmanager list-secrets

# ====================================================================================
# Golang

# go-test:
# ifeq ($(EXPORT_RESULT), true)
# 	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
# 	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
# endif
# 	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

# coverage:
# 	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
# 	$(GOCMD) tool cover -func profile.cov
# ifeq ($(EXPORT_RESULT), true)
# 	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
# 	GO111MODULE=off go get -u github.com/axw/gocov/gocov
# 	gocov convert profile.cov | gocov-xml > coverage.xml
# endif

go-lint: ## Use golintci-lint on your project
	@if ! golangci-lint run; then \
		echo -e "\033[0;33mgolangci-lint failed: some checks can be fixed with \`\033[0;32mmake fmt\033[0m\033[0;33m\`\033[0m"; \
		exit 1; \
	fi
	@$(OK) Finished linting

go-vendor: ## calls the system to download the dependencies and put them in the /vendor directory
	$(GOCMD) mod vendor


go-build: go-get ## Build binary
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)


go-get: go-vendor ## Checking if there is any missing dependencies
	@echo "  >  Checking if there is any missing dependencies..."
	$(GOCMD) get ./...
	$(GOCMD) mod tidy



go-install: go-get
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)


go-clean: ## Cleaning build cache
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


# ====================================================================================
# Help

# only comments after make target name are shown as help text
help: ## Displays this help message
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s : | sort)"
