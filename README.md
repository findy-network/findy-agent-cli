# findy-agent-cli

![Build and test](https://github.com/findy-network/findy-agent-cli/workflows/Build%20and%20test/badge.svg)

findy-agent-cli is a CLI tool for [findy-agent](https://github.com/findy-network/findy-agent) project. This tool provides some basic agency, pool & agent actions. findy-agent-cli can be used e.g. to start agency, create pool & making connections between agents.

## Environment setup

1. Install [libindy-dev](https://github.com/hyperledger/indy-sdk/#installing-the-sdk). For mac environments you may need to [install the package from sources](https://github.com/findy-network/findy-issuer-api#1-indy-sdk).
2. Install Go. Make sure environment variable `GOPATH`is defined.
3. Create parent folder for findy-agent-project in your [\$GOPATH](https://github.com/golang/go/wiki/SettingGOPATH):

   ```
   $GOPATH/src/github.com/findy-network
   ```

4. Clone [findy-agent-cli](https://github.com/findy-network/findy-agent-cli) (or move) repository to the newly created parent folder.

5. Install binary: `make install`

   If build system cannot find indy libs and headers, set following environment
   variables:

   ```
   export CGO_CFLAGS="-I/<path_to_>/indy-sdk/libindy/include"
   export CGO_LDFLAGS="-L/<path_to_>/indy-sdk/libindy/target/debug"
   ```

6. Binary should appear in `$GOPATH/bin/findy-agent-cli`

   Use --help flag after desired command to see detailed usage explanation of the command.

### Running agency

Repo contains helper scripts for setting up ledger and findy-agency. NOTE: use with caution! **Helper scripts erase all local indy-data from ~/.indy_client**

Starts test-ledger (docker), sets up steward and launches agency:

```bash
make scratch
```

After first run, agency can be restarted with

```bash
make run
```

See examples how to onboard clients to agency from [findy-agent](https://github.com/findy-network/findy-agent), [findy-issuer-api](https://github.com/findy-network/findy-issuer-api) or [findy-wallet-ios](https://github.com/findy-network/findy-wallet-ios) repositories.

Agency API documentation can be found [here](https://github.com/findy-network/findy-agent-api).

TODO: add example how to onboard client with CLI and create schema etc.

### Docker

To build CLI docker image run: `make image`

Example usage:

`docker run --network="host" --rm findy-agent-cli service ping`

note: use --network="host" flag to use host computer network settings.

## Usage

In addition to passing command flags into your shell command, it is possible to use enviroment variables or configuration files to specify your flag values.

### Configuration file

In order to use configuration file place your configuration file path to --config flag.

Example: `findy-agent-cli agency start --config path/to/my/config.yaml`

### ENV variable usage

You can pass flag values using enviroment variables.

Example: `export FCLI_AGENCY_STEWARD_SEED="findy-cli-config.yaml"`

ENV variable names can be found from flag usage info. To see flag info for specific command, use `--help` flag after the command.

Example: `findy-agent-cli agency start --help`

### Shell autocompletion

Use `findy-agent-cli completion <shell type>` command to generate findy-agent-cli autocompletion script to your shell enviroment.

You can load bash autocomletion code into your current shell with these commands:

bash: `source <(findy-agent-cli completion bash)`
zsh: `source <(findy-agent-cli completion zsh)`

Note! Bash autocompletion requires [bash-completion](https://github.com/scop/bash-completion) to be installed beforehand.

#### Enable to all shell sessions (optional)

According which shell you are using, add one of the previous commands to your shell configuration scripts (e.g. .bash_profile/.zshrc)
