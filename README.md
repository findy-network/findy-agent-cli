# findy-agent-cli

![Build and test](https://github.com/findy-network/findy-agent-cli/workflows/Build%20and%20test/badge.svg)

CLI tool for findy-agent.

## Usage



## Autocompletion

findy-cli autocompletion files can be found in scripts folder. 

For example to enable bash-completion, run bash_completion.sh script.

## Docker image

To build docker image run: `make image`

Example usage: `docker run --network="host" --rm findy-cli service`

note: use --network="host" flag  to use host computer network settings.
