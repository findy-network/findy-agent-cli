# Container image that runs your code
FROM optechlab/indy-golang:1.16.0

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY ./findy-agent-cli/action/e2e/entrypoint.sh /entrypoint.sh
COPY ./findy-agent-cli /findy-agent-cli

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
