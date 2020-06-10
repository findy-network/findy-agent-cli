ARG CLI_VERSION=latest

FROM findy-agent-cli:$CLI_VERSION

RUN echo "{}" > /root/findy.json

EXPOSE 8080

# debug levels 3 - 5 - 10 from less to more verbose
ENV FINDY_AGENT_CLI_LOGGING -logtostderr=true -v=3

ENV FINDY_AGENT_CLI_HOST_ADDR localhost
ENV FINDY_AGENT_CLI_HOST_PORT 8080
ENV FINDY_AGENT_CLI_SERVER_PORT 8080
ENV FINDY_AGENT_CLI_GENESIS /genesis_transactions
ENV FINDY_AGENT_CLI_POOL findy
ENV FINDY_AGENT_CLI_PSMDB /root/findy.bolt
ENV FINDY_AGENT_CLI_REGISTER /root/findy.json
ENV FINDY_AGENT_CLI_STEWARD_NAME steward

ENV FINDY_AGENT_CERT_PATH /aps.p12

CMD ["/start-server.sh", "/"]