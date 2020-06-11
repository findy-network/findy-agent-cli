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

issuer-api:
	docker run --network="host" --rm findy-agent-cli service onboard \
	--url http://localhost:8080 \
	--wallet-name issuer-wallet \
	--wallet-key CgM78xxAahCBG1oUrnRE3iy73ZjxbjQGuVYs2WoxpZKE \
	--email issuer-wallet-email \
	--export-filepath ~/exports/issuer-wallet \
