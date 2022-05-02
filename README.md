# findy-agent-cli

[![test](https://github.com/findy-network/findy-agent-cli/actions/workflows/test.yml/badge.svg?branch=dev)](https://github.com/findy-network/findy-agent-cli/actions/workflows/test.yml)

## Getting Started

Findy Agency is a collection of services ([Core](https://github.com/findy-network/findy-agent),
[Auth](https://github.com/findy-network/findy-agent-auth),
[Vault](https://github.com/findy-network/findy-agent-vault) and
[Web Wallet](https://github.com/findy-network/findy-wallet-pwa)) that provide
full SSI agency along with a web wallet for individuals.
To start experimenting with Findy Agency we recommend you to start with
[the documentation](https://findy-network.github.io/) and
[set up the agency to your localhost environment](https://github.com/findy-network/findy-wallet-pwa/tree/dev/tools/env#agency-setup-for-local-development).

- [Documentation](https://findy-network.github.io/)
- [Instructions for starting agency in Docker containers](https://github.com/findy-network/findy-wallet-pwa/tree/dev/tools/env#agency-setup-for-local-development)

## Project

findy-agent-cli is a command-line tool for
[Findy Agency](https://github.com/findy-network/findy-agent) Aries protocol
engine. The tool is a standalone CLI with minimal dependencies i.e. none.
The binary includes all that's needed to run it. It provides commands to

- allocate new cloud agent;
- authenticate thru WebAuthn by including headless FIDO2 authenticator which can
  be easily used from other processes thru JSON interface as well;
- communicate with the cloud agent and make connections, invitations, cred defs,
  schemas, and listen the agent notifications;
- execute protocol commands like issue credential, request proof, trust-ping,
  etc.;
- start and communicate with chat-bots;
- execute operation commands to agency itself like setting logging levels,
  querying statistics;
- and naturally auto-completion is supported

## Installation

For fast installation use the `install.sh` script. For the brave ones:

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/findy-network/findy-agent-cli/HEAD/install.sh)"
```

Or just download the script first:

```shell
curl https://raw.githubusercontent.com/findy-network/findy-agent-cli/HEAD/install.sh
```

The default installation directory is `./bin`. If you want to set it first run
the `install.sh` with the `-b` flag. Example for the macOS:

```shell
./install.sh -b /usr/local/bin
```

Don't forget to set the [auto-completion](#Shell-auto-completion)

## Running Full Stack Agency

For running quickly a full-stack trial please see the
[guide](./scripts/fullstack/README.md).
It includes all key components of the Findy Agency in docker compose file and
detailed examples how to use `findy-agent-cli` to setup wallets and cloud
agents.

## Building From Source Files

Follow these steps to install CLI tool from the source. Please make sure that Go
and git are both installed and working properly. You should have `$GOPATH/bin/`
in your $PATH variable.

_Note!_ Go modules must be on.

1. Clone [findy-agent-cli](https://github.com/findy-network/findy-agent-cli)
2. Install binary: `make install` or `make deps cli`
3. Binary will in `$GOPATH/bin/` by name `findy-agent-cli` or if you used `make deps cli` by name `cli`
4. To activate auto-completion run: `. sa-compl.sh` without arguments if you
   used `make deps cli` and use `. sa-compl.sh findy-agent-cli findy-agent-cli`
   if you used the first option.

## CLI usage examples

Examples are now included in the command help which can be activated by `cli agency -h` or `cli help agency`.

More examples can be found from `scripts/fullstack/README.md`

## Usage

In addition to passing flags into the command, it is possible to use environment
variables or configuration files to specify your flag values.

### Configuration file

In order to use configuration file place your configuration file path to
`--config` flag.

Example: `findy-agent-cli bot start --config my_config.yaml`

##### Dev Tip

If you have `export FCLI_CONFIG=./cfg.yaml` in your environment variables you
easily can have directory based configurations to execute CLI-tools commands
just by defining `cfg.yaml` files to those directories you want to present your
agent. Only thing you have to do is switch to the directory which contains the
proper `cfg.yaml` for the context you want to use. The directory name tells you
the context you are in.

### ENV variable usage

You can pass flag values using environment variables.

The example of typical and minimal settings for usage of the tool:

```
export FCLI_SERVER="host.domain.net:50051"
export FCLI_URL="https://host.domain.net"
export FCLI_TLS_PATH="/home/god/go/src/github.com/findy-network/certs"
```

The previous environment variables tell the findy-agent-cli where the WebAuthn
server and gRPC agency are as well set the proper TLS certificate path.

ENV variable names can be found from flag usage info. To see flag info for
specific command, use `--help` flag after the command.

Example: `findy-agent-cli authn --help`

### Shell auto-completion

Use `findy-agent-cli completion <shell type>` command to generate
findy-agent-cli auto-completion script to your shell environment.

You can load bash auto-completion code into your current shell with these
commands:

Bash: `source <(findy-agent-cli completion bash)`
zsh: `source <(findy-agent-cli completion zsh)`

Note! Bash auto-completion requires
[bash-completion](https://github.com/scop/bash-completion) to be installed
beforehand.

##### Dev Tip

There is `sa-compl.sh` helper script to allow easily take use of auto-completion
at the run-time even with complex aliases like this:

```shell script
alias your_alias='go run .'
. ./sh-compl.sh "go run ." your_alias
```

If you are using `make cli` to build dev-time version of the command to your
`GOPATH` you can just run `$ . ./sh-compl.sh`

#### Enable to all shell sessions (optional)

According which shell you are using, add one of the previous commands to your
shell configuration scripts (e.g. .bash_profile/.zshrc)
