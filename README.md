# findy-agent-cli

![Build and test](https://github.com/findy-network/findy-agent-cli/workflows/Build%20and%20test/badge.svg) ![e2e test](https://github.com/findy-network/findy-agent-cli/workflows/e2e%20test/badge.svg)


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

Agency API documentation can be found [here](https://github.com/findy-network/findy-agent-api).

## CLI usage examples

### Edge Agent On-boarding

Findy-agent serves all types of edge agents (EA) by implementing a corresponding
cloud agent (CA) for them. An EA communicates with its CA with Aries's
predecessor of DIDComm, which means that the communication between EA and CA
needs indy SKD's wallet and crypto functions. The same mechanism is used when
the agency calls a service agent (SA), a particular type of an EA which performs
as an issuer or verifier or both.

The agency offers an API to make a handshake, aka onboarding, where a new CA is
allocated and bound to an EA. findy-agent-cli can call that same API by itself as a
client, a temporary EA. That is an easy way to onboard SAs to the agency. The
following command is an example of calling an API to make a handshake and export
the client wallet and move it where the final SA will run.

Example:
```
  findy-agent-cli service onboard \
    --agency-url=http://localhost:8080 \
    --wallet-name=my_wallet \
	--wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY	\
	--email=my_email \
	--salt=my_salt \
    --export-file=pathto/my_wallet.export \
    --export-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY
```

### Agent to agent connection

Agents can use invitation messages to make connection to each other. Invitation message for agent can be created via invitation command.

```
  findy-agent-cli user invitation \
	--wallet-name=my_wallet \
	--wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
	--label=invitation_label
```

To use invitation file to connect, pass the file as command argument.

```
  findy-agent-cli service connect path/to/invitationFile
```

You can also read invitation json from standard input.

```
  findy-agent-cli service connect - {invitationJson}
```

To make connection without using invitation message.

Example:
```
  findy-agent-cli service connect \
    --wallet-name=my_wallet \
	--wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
	--name=my_pairwise_name \
	--key=my_pairwise_key \
	--endpoint=pairwise_endpoint
```

### Creating schema

Only service agents are able to create schemas. You need to specify name, version and attributes of the schema.

Example:
```
  findy-agent-cli service schema create \
    --wallet-name=my_wallet \
    --agency-url=http://localhost:8080 \
    --wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --name=my_schema \
    --version=2.0 \
    --attributes=["field1", "field2", "field3"]
```

### Creating credential definition

Only service agents are able to create credential definitions. You need to specify tag and schema-id of the credential definition.

Example:
```
  findy-agent-cli service creddef create \
    --wallet-name=my_wallet \
    --agency-url=http://localhost:8080 \
    --wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --schema-id=my_schema_id \
    --tag=my_creddef_tag \
```

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

## Running e2e tests

Run end-to-end tests for findy-agent-cli with:

```
make e2e
```
This starts test-ledger & runs e2e tests for findy-agent-cli.

`make e2e_ci` doesn't initialize test-ledger.

