# findy-agent-cli

![Build and test](https://github.com/findy-network/findy-agent-cli/workflows/Build%20and%20test/badge.svg)

findy-agent-cli is a CLI tool for [findy-agent](https://github.com/findy-network/findy-agent) project. This tool provides some basic agency, pool & agent actions. findy-agent-cli can be used e.g. to start agency, create pool & making connections between agents.  

## Get Started

1. [Install](https://github.com/hyperledger/indy-sdk/#installing-the-sdk) libindy-dev.
2. Clone the repo: `git clone https://github.com/findy-network/findy-agent-cli.git`
3. Install needed Go packages: `make deps`

If build system cannot find indy libs and headers, set following environment 
variables:

```
export CGO_CFLAGS="-I/<path_to_>/indy-sdk/libindy/include"
export CGO_LDFLAGS="-L/<path_to_>/indy-sdk/libindy/target/debug"
```

Use --help flag after desired command to see detailed usage explanation of the command.

### About flag usage

In addition to passing command flags into your shell command, it is possible to use enviroment variables or configuration files to specify your flag values.

In order to use configuration file place your configuration file path to --config flag.

Example: `findy-agent-cli agency start --config path/to/my/config.yaml`

You can pass flag values using enviroment variables by adding `FINDY_AGENT_CLI` in front of your flag name.

Example: `export FINDY_AGENT_CLI_PSMDB="findy-cli.bolt"`

Note: when using enviroment variables, you must replace `-` character with `_`.

Example: `export FINDY_AGENT_CLI_WALLET_NAME="testWallet"` 

## Shell autocompletion

Use `findy-agent-cli completion <shell type>` command to generate findy-agent-cli autocompletion script to your shell enviroment.

You can load bash autocomletion code into your current shell with these commands:

bash: `source <(findy-agent-cli completion bash)`
zsh: `source <(findy-agent-cli completion zsh)`

Note! Bash autocompletion requires [bash-completion](https://github.com/scop/bash-completion) to be installed beforehand.

#### Enable to all shell sessions (optional)

According which shell you are using, add one of the previous commands to your shell configuration scripts (e.g. .bash_profile/.zshrc) 

## Docker usage

To build docker image run: `make image`

Example usage: `docker run --network="host" --rm findy-agent-cli service ping`

note: use --network="host" flag to use host computer network settings.
