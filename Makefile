AGENT_PATH=github.com/findy-network/findy-agent
LEDGER_NAME:=FINDY_FILE_LEDGER

AGENT_BRANCH=$(shell ./branch.sh ../findy-agent/)
API_BRANCH=$(shell ./branch.sh ../findy-agent-api/)
AUTH_BRANCH=$(shell ./branch.sh ../findy-agent-auth/)
GRPC_BRANCH=$(shell ./branch.sh ../findy-common-go/)
WRAP_BRANCH=$(shell ./branch.sh ../findy-wrapper-go/)

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
	@./scan.sh

build:
	go build ./...

cli:
	go build -o $(GOPATH)/bin/cli

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

test_cov:
	go test -v -p 1 -failfast -coverprofile=c.out ./... && go tool cover -html=c.out

check: check_fmt vet shadow

install:
	$(eval VERSION = $(shell cat ./VERSION))
	@echo "Installing version $(VERSION)"
	go install \
		-ldflags "-X '$(AGENT_PATH)-cli/utils.Version=$(VERSION)'" \
		./...

# https://goreleaser.com/install/
test_release:
	goreleaser --snapshot --skip-publish --rm-dist

