# Agency Setup for Local Development

## CLI Agents

**Note!** The repo's `.gitignore` already defines few directory prefix to use
your own sandboxes: `op*`, `my*`, `new*`, and `test*`

You can build your own playground agents by running the following in the command
line:

```console
make-agent.sh <agent-name>
```

## Setup

In case you do not have cloud installation of Findy Agency available, you can
setup needed services locally and develop your application against a local Findy
Agency. This document describes how to set up Findy Agency service containers on
your local computer.

The setup uses agency internal file ledger, intended only for testing during
development. This setup does not suit for testing inter-agency communication
even though it is possible to set one up using a common indy-plenum ledger.

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [findy-agent-installation](https://github.com/findy-network/findy-agent-cli#installation)

## Steps

1. Launch backend services with

   ```sh
   make pull-up
   ```

   This will pull the latest versions of the needed docker images. Later on when
   launching the backend you can use `make up` if there is no need to fetch the
   latest images.

   It will take a short while for all the services to start up. Logs from all of
   the started services are printed to the console. `<CTRL>+C` stops the
   containers.

   The script will create a folder called `.data` where all the data of the
   services are stored during execution. If there is no need for the test data
   anymore, `make clean` will remove all the generated data and allocated
   resources.

   **UPDATE:** because of PostgreSQL and docker-compose problems in macOS the
   default `docker-compose.yml` file handles SQL db files as transient. To make
   them persistent, *remove* following line from the compose file:
   ```shell
      PGDATA: /var/lib/pg_data # <- remove me to keep vault files persistent
   ```

1. Build playground environment with CLI tool. It's usually
   good idea to have some test data at the backend before UI development or
   application logic itself. Now, when your whole stack is running thanks to
   step one you can easily play with it from the command line.

   To install `findy-agent-cli` execute the following:
   ```console
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/findy-network/findy-agent-cli/HEAD/install.sh)"
   ```
   It will install the one binary which is only that's needed in
   `./bin/findy-agent-cli` where you can move it to your path, create alias for
   it, setup auto-completion, etc. More information about it can be found from
   [here](https://github.com/findy-network/findy-agent-cli#installation).

   You should enter the following after you have installed the working
   `findy-agent-cli`:
   ```console
   export FCLI=<your-name-for-binary>
   ```
   That's for the helper scrips used in this directory and referenced here as
   well.

   To make use of `findy-agent-cli` there is a helper script to setup the CLI
   environment. Enter the following command:
   ```console
   source ./setup-cli-env.sh 
   ```
   That will setup all the needed environment variables for CLI configuration
   for the currently running environment. Most importantly it creates a new
   master key for your CLI FIDO2 authenticator. If you want to keep your
   development environment between restarts you should persist that key by
   copying it to your environment variables. The key is in env `FCLI_KEY` after
   running the setup script. The setup generates the `use-key.sh` script for
   your convenient as well. Next time you source `./setup-cli-env.sh` it uses
   `use-key.sh` if it exists. If you want to create a new key you should remove
   `use-key.sh` before sourcing `./setup-cli-env.sh`.

   *Tip* Enter following commands:
   ```console
   alias cli=findy-agent-cli 
   . <(findy-agent-cli completion bash | sed 's/findy-agent-cli/cli/g')
   ```

   *Tip For Linux only*: 

   - define following aliases and install `xclip` if not
     already installed:
     ```sh
     alias pbcopy='xclip -selection c'
     alias pbpaste='xclip -selection clipboard -o'
     alias pbtee='pbcopy && pbpaste'
     ```

   **On-board Alice and Bob and Government**
   ```console
   alice/register
   bob/register
   government/register
   ```
   You can play each of them by entering for example following:
   ```console
   alice/login
   pushd alice
   $FCLI agent ping
   popd
   ```

   **Alice invites Bob to connect**

   Enter following commands:
   ```console
   alice/invitation | bob/connect | pbtee
   ```

   Now you have the connection ID (pairwise ID) in the environment variable and
   in your clipboard. You can test connection with the commands:
   ```console
   cd alice/$(pbpaste)
   $FCLI connection trustping
   ```
   Which means that Alice's end of the connection calls Aries's trustping
   protocol and Bob's cloud agent responses it.

   Before entering previous commands you could open *a second terminal window*
   and execute following:
   ```console
   source ./setup-cli-env.sh
   cd bob/$(pbpaste)
   $FCLI agent listen
   ```
   You should now start to receive notifications of the Aries protocols. You
   also see those notifications which were triggered before `listen`.

   **Alice sends text message to Bob**

   First in the Bob's terminal stop the previous listening with C-c and enter
   the following:
   ```console
   $FCLI bot read
   ```
   Go to the Alice's terminal and enter the commands:
   ```console
   # You should be in directory alice/<connection-id>
   echo 'Hello Bob! Alice here.' | $FCLI bot chat
   ```
   The Bob's terminal should output Alice's welcoming messages.
   
   Congratulations! You have made this far and successfully run two Aries
   protocols `trustping` and `basic message`. Next we will cover two most
   important ones `issue credential` and `present proof`.

   **Government creates Schema and CredDef**

   First switch Alice's terminal and login as Government.
   ```console
   government/login
   ```
   As Government create a new schema and a credential definition.
   ```console
   source government/new-schema
   source government/new-cred-def  # please be patient, this will take time
   ```
   **Government invites Bob to connect**

   Before we can issue the newly created credential to Bob we must make a
   connection between Government and Bob.
   ```console
   government/invitation | bob/connect | pbtee
   ```
   Go to back to Bob's terminal and stop `bot read` command with C-c and enter
   the following:
   ```console
   # make sure you are bob/ directory
   cd $(pbpaste)
   $FCLI agent mode-cmd --auto   # this is very important!
   $FCLI agent listen
   ```
   You might wonder what for the command `$FCLI agent mode-cmd --auto` is given.
   It sets Bob's cloud agent to automatically accept all issued credentials or
   even automatically make a proof when requested. Normally that would need a
   permission from the controller, the wallet owner. The auto-mode is especially
   useful for tests, samples, etc.

   **Government trustpings Bob**
   
   Go to Government terminal and enter following to verify newly created DIDComm
   connection:
   ```console
   cd $(pbpaste)
   $FCLI connection trustping
   ```
   You should see the notification in Bob's terminal and get OK status here.

   **Government issues credential to Bob**

   Run the following helper script:
   ```console
   ../issue bob
   ```
   You should get OK indication as a command return and you should see the
   notification on Bob's terminal that a new credential is received.

   **Government requerst proof from Bob**

   First we must set Government into auto-accept mode because when full proof
   protocol is running it must verify data when controller is called. By setting
   the auto-mode the command do it for us:
   ```console
   $FCLI agent mode-cmd --auto   # this is very important!
   ```

   Enter the following to request a proof from Bob:
   ```console
   government/request-proof
   ```
   You should get OK indication as a command return and you should see the
   notification on Bob's terminal that proof is made.

   Congratulations! Now you have run all the most important DIDComm protocols by
   your own. More samples and guides can be found from
   [Findy Wallet](https://github.com/findy-network/findy-wallet-pwa).

   **Admin Operations**

   After environment setup you can see what your configuration is by executing
   the following helper script:
   ```console
   admin/cli-env
   ```
   It will output all of the `findy-agent-cli` env configurations currently set.
   To check one specific variable enter: `admin/cli-env KEY` for example.

   To register your CLI authenticator for direct communication to Findy Agency
   enter the following commands:
   ```console
   source admin/register
   source admin/login
   ```
   Later the login is all what is needed. After successful login you can enter
   commands like:
   ```console
   $FCLI agency count         # get status of the clould agents
   $FCLI agency logging -L=5  # set login level of the core agency 
   ```

# Problem Solving Dev and Demo Environment

## CLI FIDO2 Authenticator and Origin

The command-line FIDO2 is a tricky one especially when working with
non-production environment where TLS termination and origin can be set
differently. If you are switching you CLI usage between public cloud agency
installation and your local installation, you have to be extra careful with our
configurations and environment variables. The following is from the local setup:

```sh
# FIDO2 server `findy-agent-auth` address
export FCLI_URL=http://localhost:8088
# Set the origin according to where our Web Wallet is hosted **important**
export FCLI_ORIGIN=http://localhost:3000
```

When a public cloud uses a reverse proxy and a load balancer you don't need to
set `FCLI_ORIGIN` or more importantly you **must not set it**.

## Authenticator's Master Key

If you continue working with existing Findy agency installation you must be
careful your CLI authenticator's master key. For that look for `use-key.sh`
file.

## Helper Scripts And Execution Environment

Don't use `--logging` argument or more precisely don't use its env version
(FCLI_CLI_LOGGING) or config file version. The scripts aren't optimised to handle
different streams for err and std. However, if it's mandatory for your use case,
read `glog` documentation how to use files for logs.

The helper scripts should be used only as an examples. They don't support all
the possible argument combinations.
