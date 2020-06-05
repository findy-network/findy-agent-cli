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

image:
	-git clone git@github.com:findy-network/findy-wrapper-go.git .docker/findy-wrapper-go
	-git clone git@github.com:findy-network/findy-agent.git .docker/findy-agent
	docker build -t findy-agent-cli .

issuer-api:
	docker run --network="host" --rm findy-cli service onboard \
	--url http://localhost:8080 \
	--walletname issuer-wallet \
	--walletkey CgM78xxAahCBG1oUrnRE3iy73ZjxbjQGuVYs2WoxpZKE \
	--cloudpw issuer-wallet \
	--email issuer-wallet \
	--exportPath ~/exports/ \
	--export=true
