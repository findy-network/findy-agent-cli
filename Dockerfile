FROM optechlab/indy-golang:1.15.0

ARG HTTPS_PREFIX

WORKDIR /findy-agent-cli

RUN git config --global url."https://"${HTTPS_PREFIX}"github.com/".insteadOf "https://github.com/"

COPY go* ./
RUN go mod download

COPY . ./
RUN make install

FROM optechlab/indy-base:1.15.0

COPY --from=0 /go/bin/findy-agent-cli /findy-agent-cli

ENTRYPOINT ["/findy-agent-cli"]
