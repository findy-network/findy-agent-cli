ARG CLI_VERSION=latest

FROM findy-agent-cli:$CLI_VERSION

ADD ./start-server.sh /start-server.sh

RUN echo "{}" > /root/findy.json

EXPOSE 8080

# debug levels 3 - 5 - 10 from less to more verbose
ENV FINDY_AGENT_CLI_LOGGING -logtostderr=true -v=3

ENV FINDY_AGENT_CLI_HOST_ADDRESS localhost
ENV FINDY_AGENT_CLI_HOST_PORT 8080
ENV FINDY_AGENT_CLI_SERVER_PORT 8080
ENV FINDY_AGENT_CLI_POOL_NAME findy
ENV FINDY_AGENT_CLI_PSM_DATABASE_FILE /root/findy.bolt
ENV FINDY_AGENT_CLI_REGISTER_FILE /root/findy.json
ENV FINDY_AGENT_CLI_STEWARD_NAME steward

ENV FINDY_AGENT_CLI_APNS_P12_FILE /aps.p12

ENTRYPOINT [ "/start-server.sh", "/"]