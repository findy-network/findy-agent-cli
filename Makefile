AGENT_PATH=github.com/findy-network/findy-agent
LEDGER_NAME:=FINDY_FILE_LEDGER

AGENT_BRANCH=$(shell ./branch.sh ../findy-agent/)
API_BRANCH=$(shell ./branch.sh ../findy-agent-api/)
GRPC_BRANCH=$(shell ./branch.sh ../findy-grpc/)

modules:
	@echo Syncing modules for work brances ...
	go get github.com/findy-network/findy-agent-api@$(API_BRANCH)
	go get github.com/findy-network/findy-grpc@$(GRPC_BRANCH)
	go get github.com/findy-network/findy-agent@$(AGENT_BRANCH)

deps:
	go get -t ./...

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
	$(GOPATH)/bin/golint ./...

lint_e:
	@$(GOPATH)/bin/golint ./... | grep -v export | cat

test:
	go test -v -p 1 -failfast ./...

test_cov:
	go test -v -p 1 -failfast -coverprofile=c.out ./... && go tool cover -html=c.out

e2e: install
	./scripts/dev/e2e-test.sh init_ledger
	./scripts/dev/e2e-test.sh e2e
	./scripts/dev/e2e-test.sh clean

e2e_ci: install
	./scripts/dev/e2e-test.sh e2e

check: check_fmt vet shadow

install:
	$(eval VERSION = $(shell cat ./VERSION))
	@echo "Installing version $(VERSION)"
	go install \
		-ldflags "-X '$(AGENT_PATH)-cli/utils.Version=$(VERSION)' -X '$(AGENT_PATH)/agent/utils.Version=$(VERSION)'" \
		./...

clean:
	-rm -rf .docker

image:
	$(eval VERSION = $(shell cat ./VERSION))
	-git clone git@github.com:findy-network/findy-wrapper-go.git .docker/findy-wrapper-go
	-git clone git@github.com:findy-network/findy-agent.git .docker/findy-agent
	docker build -t findy-agent-cli .
	docker tag findy-agent-cli:latest findy-agent-cli:$(VERSION)

agency: image
	$(eval VERSION = $(shell cat ./VERSION))
	docker build -t findy-agency --build-arg CLI_VERSION=$(VERSION) ./agency
	docker tag findy-agency:latest findy-agency:$(VERSION)

# Test for agency-image start script:
#run-agency: agency
#	echo "{}" > findy.json && \
#	docker run -it --rm -v $(PWD)/agency/infra/.secrets/steward.exported:/steward.exported \
#		-e FCLI_AGENCY_SALT="this is only example" \
#		-p 8080:8080 \
#		-v $(PWD)/agency/infra/.secrets/aps.p12:/aps.p12 \
#		-v $(PWD)/scripts/dev/genesis_transactions:/genesis_transactions \
#		-v $(PWD)/findy.json:/root/findy.json findy-agency

# **** scripts for local agency development:
# WARNING: this will erase all your local indy wallets
scratch:
	./scripts/dev/dev.sh scratch $(LEDGER_NAME)

run:
	./scripts/dev/dev.sh install_run $(LEDGER_NAME)
# ****
