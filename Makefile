AGENT_BRANCH=$(shell scripts/branch.sh ../findy-agent/)
API_BRANCH=$(shell scripts/branch.sh ../findy-agent-api/)
AUTH_BRANCH=$(shell scripts/branch.sh ../findy-agent-auth/)
GRPC_BRANCH=$(shell scripts/branch.sh ../findy-common-go/)
WRAP_BRANCH=$(shell scripts/branch.sh ../findy-wrapper-go/)

SCAN_SCRIPT_URL="https://raw.githubusercontent.com/findy-network/setup-go-action/master/scanner/cp_scan.sh"

cli:
	$(eval VERSION = $(shell cat ./VERSION) $(shell date))
	@echo "Installing version $(VERSION)"
	go build \
		-ldflags "-X 'github.com/findy-network/findy-agent-cli/utils.Version=$(VERSION)'" \
		-o $(GOPATH)/bin/cli

drop_wrap:
	go mod edit -dropreplace github.com/findy-network/findy-wrapper-go

drop_comm:
	go mod edit -dropreplace github.com/findy-network/findy-common-go

drop_auth:
	go mod edit -dropreplace github.com/findy-network/findy-agent-auth

drop_api:
	go mod edit -dropreplace github.com/findy-network/findy-agent-api

drop_agent:
	go mod edit -dropreplace github.com/findy-network/findy-agent

drop_all: drop_api drop_comm drop_wrap drop_wrap drop_auth

repl_wrap:
	go mod edit -replace github.com/findy-network/findy-wrapper-go=../findy-wrapper-go

repl_comm:
	go mod edit -replace github.com/findy-network/findy-common-go=../findy-common-go

repl_api:
	go mod edit -replace github.com/findy-network/findy-agent-api=../findy-agent-api

repl_auth:
	go mod edit -replace github.com/findy-network/findy-agent-auth=../findy-agent-auth

repl_agent:
	go mod edit -replace github.com/findy-network/findy-agent=../findy-agent

repl_all: repl_api repl_comm repl_wrap repl_agent repl_auth

modules: modules_api modules_auth modules_wrap modules_comm modules_agent

modules_api: drop_api
	@echo Syncing modules: findy-agent-api/$(API_BRANCH)
	go get github.com/findy-network/findy-agent-api@$(API_BRANCH)

modules_auth: drop_auth
	@echo Syncing modules: findy-agent-api/@$(AUTH_BRANCH)
	go get github.com/findy-network/findy-agent-auth@$(AUTH_BRANCH)

modules_wrap: drop_wrap
	@echo Syncing modules: findy-agent-api/$(WRAP_BRANCH)
	go get github.com/findy-network/findy-wrapper-go@$(WRAP_BRANCH)

modules_comm: drop_comm
	@echo Syncing modules: findy-agent-api/$(GRPC_BRANCH) 
	go get github.com/findy-network/findy-common-go@$(GRPC_BRANCH)

modules_agent: drop_agent
	@echo Syncing modules: findy-agent-api/$(AGENT_BRANCH) 
	go get github.com/findy-network/findy-agent@$(AGENT_BRANCH)

deps:
	go get -t ./...

scan:
	@curl -s $(SCAN_SCRIPT_URL) | bash

scan_and_report:
	@curl -s $(SCAN_SCRIPT_URL) | bash -s v > licenses.txt

build:
	go build ./...

misspell:
	@go get github.com/client9/misspell 
	@find . -name '*.md' -o -name '*.go' -o -name '*.puml' | xargs \
		misspell -error -locale GB

vet:
	go vet ./...

shadow:
	@echo Running govet
	go vet -vettool=$(GOPATH)/bin/shadow ./...
	@echo Govet success

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

lint_e:
	@$(GOPATH)/bin/golint ./... | grep -v export | cat

test:
	go test -v -p 1 -failfast ./...

test_cov_out:
	go test \
		-coverpkg=github.com/findy-network/findy-agent-cli/... \
		-coverprofile=coverage.txt  \
		-covermode=atomic \
		./...

test_cov: test_cov_out
	go tool cover -html=coverage.txt

check: check_fmt vet shadow

install:
	$(eval VERSION = $(shell cat ./VERSION) $(shell date))
	@echo "Installing version $(VERSION)"
	go install \
		-ldflags "-X 'github.com/findy-network/findy-agent-cli/utils.Version=$(VERSION)'" \
		./...

# https://goreleaser.com/install/
test_release:
	goreleaser --snapshot --skip-publish --rm-dist

release:
	gh workflow run do-release.yml

