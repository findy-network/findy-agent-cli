deps:
	go get -t ./...

build:
	go build ./...

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

check: check_fmt vet shadow

install:
	$(eval VERSION = $(shell cat ./VERSION))
	@echo "Installing version $(VERSION)"
	go install \
		-ldflags "-X 'github.com/findy-network/findy-agent-cli/utils.Version=$(VERSION)'" \
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

clean:
	rm -rf .docker


#run-agency: agency
#	echo "{}" > findy.json && \
#	docker run -it --rm -v $(PWD)/steward.exported:/steward.exported \
#		-v $(PWD)/genesis_transactions:/genesis_transactions \
#		-v $(PWD)/findy.json:/root/findy.json findy-agency

issuer-api:
	docker run --network="host" --rm findy-agent-cli service onboard \
	--agency-url http://localhost:8080 \
	--wallet-name issuer-wallet \
	--wallet-key CgM78xxAahCBG1oUrnRE3iy73ZjxbjQGuVYs2WoxpZKE \
	--email issuer-wallet-email \
	--export-file ~/exports/issuer-wallet \
	--export-key CgM78xxAahCBG1oUrnRE3iy73ZjxbjQGuVYs2WoxpZKE \

# **** scripts for local agency development:
# WARNING: this will erase all your local indy wallets
scratch:
	./scripts/dev/dev.sh scratch

run:
	./scripts/dev/dev.sh install_run
# ****
