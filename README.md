# findy-agent-cli

![lint](https://github.com/findy-network/findy-agent-cli/workflows/golangci-lint/badge.svg?branch=dev)
![test](https://github.com/findy-network/findy-agent-cli/workflows/test/badge.svg?branch=dev)

findy-agent-cli is a CLI tool for [findy-agent](https://github.com/findy-network/findy-agent) project. This tool provides some basic agency, pool & agent actions. findy-agent-cli can be used e.g. to start agency, create pool & making connections between agents.

When [indy-cli](https://github.com/hyperledger/indy-sdk/tree/master/cli) starts new prompt where you give commands **findy-agent-cli** is build the way that you don't have to do that. You can stay on your favorite shell and execute commands from there. findy-agent-cli includes many features to make its usage as convenient as possible.

- **performance**, everything is execute at once, and long-lasting commands give progress information.
- **autocompletion**, every command support autocompletion, even wallets can be autocompleted.
- **scriptable**, commands can be used according to Unix's famous _pipe-and-filter_ principle.
- **env and config file support**, every command support all three ways to enter command flags: command line, environment variable, and config YAML.
- **export CLI_CONFIG=your-name-here.yaml**, even without 3rd party tools findy-agent-cli can support directory based context switch.

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

Starts test-ledger (file), sets up steward and launches agency:

```bash
make scratch
```

After first run, agency can be restarted with

```bash
make run
```

If you want to use actual indy-node (in docker) for local test ledger, give ledger name as argument:

```bash
make scratch LEDGER_NAME="findy"
make run LEDGER_NAME="findy"
```

Agency API documentation can be found [here](https://github.com/findy-network/findy-agent-api).

## CLI usage examples

### Edge Agent On-boarding

Findy-agent serves all types of edge agents (EA) by implementing a corresponding
cloud agent (CA) for them. An EA communicates with its CA with Aries's
predecessor of DIDComm, which means that the communication between EA and CA
needs indy SDK's wallet and crypto functions. The same mechanism is used when
the agency calls a service agent (SA), a particular type of an EA which performs
as an issuer or verifier or both.

The agency offers an API to make a handshake, aka on-boarding, where a new CA is
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

In addition to passing command flags into your shell command, it is possible to use environment variables or configuration files to specify your flag values.

### Configuration file

In order to use configuration file place your configuration file path to --config flag.

Example: `findy-agent-cli agency start --config path/to/my/config.yaml`

##### Dev Tip

If you have `export FCLI_CONFIG=./cfg.yaml` in your environment variables you easily can have directory based configurations to execute CLI-tools commands just by defining `cfg.yaml` files to those directories you want to present your agent. Only thing you have to do is switch to the directory which contains the proper `cfg.yaml` for the context you want to use. The directory name tells you the context you are in.

### ENV variable usage

You can pass flag values using environment variables.

Example: `export FCLI_AGENCY_STEWARD_SEED="findy-cli-config.yaml"`

ENV variable names can be found from flag usage info. To see flag info for specific command, use `--help` flag after the command.

Example: `findy-agent-cli agency start --help`

### Shell autocompletion

Use `findy-agent-cli completion <shell type>` command to generate findy-agent-cli autocompletion script to your shell environment.

You can load bash autocompletion code into your current shell with these commands:

bash: `source <(findy-agent-cli completion bash)`
zsh: `source <(findy-agent-cli completion zsh)`

Note! Bash autocompletion requires [bash-completion](https://github.com/scop/bash-completion) to be installed beforehand.

##### Dev Tip

There is `sa-compl.sh` helper script to allow easily take use of autocompletion at the run-time even with complex aliases like this:

```shell script
$ alias your_alias='go run .'
$ . ./sh-compl.sh "go run ." your_alias
```

If you are using `make cli` to build dev-time version of the command to your `GOPATH` you can just run `$ . ./sh-compl.sh`

#### Enable to all shell sessions (optional)

According which shell you are using, add one of the previous commands to your shell configuration scripts (e.g. .bash_profile/.zshrc)

## Using Example Scripts as A Playground

The `scripts` folder includes `test` directory to start two different types of the agencies. The `mem-server` start agency with inmemory ledger that vanish when server stops, and `file-server` which saves all ledger transactions to `FINDY_FILE_LEDGRE.json` (default `~/.indy_client/`). The file ledger allows you to run long-running agency services on your own machine without running the real ledger on the same machine.

The `scripts/clientN` folders include examples of configuration files to run _holder agents_. The example of the _server agent_ (SA) for these agents is in the `/scrips/sa` folder. Please note that you cannot directly use the configuration files but you must first onboard your own agents (allocate cloud agents from the current agency) and write the configuration files after that according to the information from onboardings, invitations, connections, etc.

#### Create Schema and Cred Def and Issue a Cred

The `script/sa` includes helper scripts `create.sh` and `issue.sh`. The first one is to create a new schema to ledger and a cred def to issuers wallet. The second one is to issue previously created cred def to other agent. Please see the `script/clientN` folders. You must make connection between e.g. SA <-> Client1 and store that information to their `cfg.yaml` file. Please follow the existing example files in their folder.

#### JWT and gRPC Version of Agency API

After on-boarding a client and allocating the cloud agent you can use JWT-based authentication systems for the agency's gRPC API access. The commands that are using the new API are sub commands of the `JWT`. Please see more information from its usage.

## Publishing new version

Release script will tag the current version and push the tag to remote. This will trigger e2e-tests in CI automatically and if they succeed, the tag is merged to master.

Release script assumes it is triggered from dev branch. It takes one parameter, the next working version. E.g. if current working version is 0.1.0, following will release version 0.1.0 and update working version to 0.2.0.

```bash
git checkout dev
./release 0.2.0
```
