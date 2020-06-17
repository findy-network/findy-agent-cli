FROM optechlab/indy-golang:1.15.0

WORKDIR /go/src/github.com/findy-network/findy-agent-cli

COPY .docker/findy-wrapper-go /go/src/github.com/findy-network/findy-wrapper-go
COPY .docker/findy-agent /go/src/github.com/findy-network/findy-agent

COPY . .

RUN make deps && make install

FROM optechlab/indy-base:1.15.0

COPY --from=0 /go/bin/findy-agent-cli /findy-agent-cli

CMD ["/findy-agent-cli"]
