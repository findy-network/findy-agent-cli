FROM optechlab/indy-golang:1.14.2

WORKDIR /go/src/github.com/optechlab/findy-cli

COPY .docker/findy-go /go/src/github.com/optechlab/findy-go
COPY .docker/findy-agent /go/src/github.com/optechlab/findy-agent

COPY . .

RUN go get -t ./... && go install

FROM optechlab/indy-base:1.14.2

COPY --from=0 /go/bin/findy-cli /findy-cli

ENTRYPOINT ["/findy-cli"]
