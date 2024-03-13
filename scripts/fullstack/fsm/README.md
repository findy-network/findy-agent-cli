# Readme

The document tells how to start with
`issuing-service-?-fsm.ymal` chatbot files. The `issuing chatbot` in the FSM
YAML files are a reference implementation meant to be customized to the
needs of your project.

### What Problem Does It Solve?

The Hyperledger Indy-based SSI system is implemented with a CL signature scheme
for ZKPs. That system needs the concept of *Credential Definition* stored in its
creator's wallet. The Credential Definition has an ID, which is quite similar to
DID in the Indy-based AnonCreds system. The CredDefID is public. Everyone who
knows it can request proof based on it or request to present a proof based on
it. 

But the CredDefID brings us some problems:
1. How do we find a correct CredDefID when it's needed?
1. How about fully symmetric cases when everyone can simultaneously be an
   issuer, a holder, and a verifier? Everyone can issue credentials and receive
   proof of them in these use cases. For example, we have a use case where a
   seller (anybody in a marketplace) wants to issue a receipt for the
   transaction.

There are other problems, but the rest are based on the same core problem.

We'll solve problem number 2 by using a notary-type service. We started with one
service and implemented a reference chatbot to issue credentials on behalf of a
logical issuer, aka seller. We also have implemented our version of a *public*
DID. With these tools, we have solved the problem quite elegantly.

#### Isn't This Centralization?

In some ways, yes, but the result isn't more centralized than the suggested
trust registries for other or similar problems in the SSI field. In a certain
way, this model adds self-sovereignty because now everyone can issue, and
everyone can build these issuing services for their use cases.

### More Reading

Before we continue, here's a list of
documents and places that are helpful when playing with these:

1. [Getting Started](https://findy-network.github.io/blog/2023/01/30/getting-started-with-ssi-service-agent-development/)
1. [Writing SSI Apps](https://findy-network.github.io/blog/2023/02/06/how-to-equip-your-app-with-vc-superpowers/)
1. [FSM documentation part I](https://findy-network.github.io/blog/2023/03/13/no-code-ssi-chatbots-fsm-part-i/)
1. [FSM documentation part II](https://findy-network.github.io/blog/2023/06/22/beautiful-state-machines-fsm-part-ii/)
1. [The workshop material for CLI tool](https://github.com/findy-network/agency-workshop/tree/master/track1-cli#task-0-setup-environment)

## Prerequisites

Note, for now, we assume that the reader uses the local setup of the agency.
It makes it easier to follow the core agent's real-time logs.

Even though you have probably cloned the repo. But if not, please do it now:

```shell
git clone https://github.com/findy-network/findy-agent-cli
```

Follow the material in the repo's readme files or the workshop material
mentioned at the beginning to get your `findy-agent-cli` CLI tool working, i.e.,
communicating with your agency.

### Helpers

1. Go to repo's root:
1. Shorter name and autocompletion:
   ```shell
   alias cli=findy-agent-cli
   source ./scripts/sa-compl.sh
   ```

> Document assumes that CLI tool is named to `cli`.

### Very Important

If you create new wallets directly with the CLI tool, make sure that auto-accept
mode is ON for the agent. The check it **after** the `cli agent ping` works:

```shell
cli agent mode-cmd -r
```

> Tip, if you have problems with CLI commands check your `FCLI_` prefixed envs.

The result should be AUTO_ACCEPT. Note that allocation scripts do this automatically.

> Chatbots should work even when auto-accept isn't ON.

# Setup Agents And Wallets

It would be best to have at least three wallets and their agents up and running.
Inside the `findy-agent-cli` repo, you have `scrpits/fullstack` directory. Let's
name it a tool root for the rest of the documentation. For example:

```shell
export MY_ROOT=`pwd`
```

Follow the workshop documentation on how to allocate new agents with their wallets.

Allocate the following agents:
1. `issuing` will be the issuing service
1. `seller`, will be a seller, aka logical issuer
1. `buyer` will be a buyer.
1. `verifier` will be a verifier. (Usage is out scoped from this document.)

Here's an example of how you could
1. allocate the agents,
1. have access to FSM files by making links,
1. create schema and credential definition (done in `issuing`),
1. create DIDComm connections between `issuing`, `seller`, and `buyer` *where
   the last is done in this phase just for testing the environment*. During the
   implementation to the app invitation or command to connect with <SessionID>
   is sent in its own step during the service use.

```shell
cd "$MY_ROOT"
make-play-agent.sh issuing seller buyer verifier
cd play/issuing
ln -s ../../fsm/issuing-service-f-fsm.yaml
ln -s ../../fsm/issuing-service-b-fsm.yaml
open `cli bot uml issuing-service-f-fsm.yaml` # UML rendering
source ./new-schema
source ./new-cred-def
./invitation | ../seller/connect
./invitation | ../buyer/connect
```

Optionally store *a public DID* of the *Issuing Service Chatbot*:

```shell
export PUB_DID=$(./pub-did print)
```

> **Note! Leave this terminal open and do not enter new commands to it yet.**

> **Note! `source ./new-cred-def` initializes `FCLI_CRED_DEF_ID` environment
> variable. The `issuing-service-f-fsm.yaml` file references to this variable,
> i.e. it's mandatory, *or* you could hard-code the *credential definition
> value* to your `issuing-service-f-fsm.yaml`.**

## Use The Issuing Service

1. open 2 separated terminals (see the workshop material on how to init envs) to
   work as `seller`, leave it to be.
1. open 2 separated terminals (see the workshop material on how to init envs) to
   work as `buyer`, leave it to be.
1. go back to the previous `issuing` terminal and start the chatbot:
   ```shell
   cli bot start --service-fsm issuing-service-b-fsm.yaml issuing-service-f-fsm.yaml -v=1
   ```
1. go back to `seller` terminal A (`cli bot read`) read-only to responses
1. go back to `seller` terminal B (`cli bot chat`)
    1. optional: enter 'help' to get used to
    1. enter <sessionID>, select something easy
    1. enter <role>, == 'issuer'
    1. enter <attr_data>, select something easy
1. go back to `buyer` terminal A (`cli bot read `) read-only to responses
1. go back to `buyer` terminal B (`cli bot chat`)
    1. optional: enter 'help' to get used to
    1. enter <sessionID>, select something easy
    1. enter <role>, == 'rcvr'
1. see both A (bot read) terminals; the results should be that the credential is
   issued.
1. go to terminal B and enter something to move FSM instances to the start state.
1. It's optional; you could rerun it with the same players.

## Sequence Diagram

Notes about current implementation:
- only one attribute value schema is implemented. Start with that and add
  cases where more attributes can be entered later.
- every message sends a `basic_message` reply, which usually starts with `ACK`. See
  the `YAML` file for more information. The reply messages aren't drawn to the
  sequence diagram below to keep it as simple as possible.
- you can render state machines to UML:
  ```shell
  open `cli bot uml issuing-service-f-fsm.yaml` # give correct FSM file
  ```

```mermaid
sequenceDiagram
    autonumber

    participant Seller

    box transparent Gray
    participant IssuerFSM
    participant BackendFSM
    participant RcvrFSM
    end

    participant Buyer

    Seller -) IssuerFSM: <session_id> (GUID)
    Seller -) IssuerFSM: issuer [= role]
    loop Schema's attributes
    Seller -) IssuerFSM: <attribute_value>
    end

    alt Send thru existing connection 
    Seller -) Buyer: <session_id> (same as above, design how app knows that this is a command)
    end

    Buyer -) RcvrFSM: <session_id>
    Buyer -) RcvrFSM: rcvr [= role]

    RcvrFSM -) BackendFSM: receiver_arriwed
    BackendFSM -) IssuerFSM: rcvr_arriwed
    loop Schema's attributes
    IssuerFSM -) BackendFSM: <attribute_value>
    BackendFSM -) RcvrFSM: <attribute_value>
    end
    IssuerFSM -) BackendFSM: attributes done (not implemented, one attrib)
    BackendFSM -) RcvrFSM: attributes done (not implemented, one attrib)

    RcvrFSM -) Buyer: CREDENTIAL ISSUING PROTOCOL
```

#### Pre-steps (not in the diagram)

1. We can generate a public DID for the *Issuing Service Chatbot*. 
   ```shell
   cd play/issuing # or where your bot is
   export PUB_DID=$(./pub-did print)
   ```
1. Save this `PUB_DID` to your app's configuration. It's where *Issuing Service
   Chatbot* can be found when needed. Note, `PUB_DID` is a URL which returns a
   new invitation on every load. You can treat is as a URL template:
   ```
   http://localhost:8080/dyn?did=8NCqkhnNjeTwMmK68DHDPx&label=<you_ad_number>
   ```
   You can enter your case specific data like ad number to the `label` arg.

#### Steps

1. Actual *Seller* role or app implementation for the role generates a sessionID 
   (GUID) and sends it to *Issuing Service Chatbot* as a `basic_message`.
1. The *Seller* role is a logical issuer, so it sends the `issuer` string as a
   `basic_message` to the *Issuer FMS* instance.
1. The *Seller* role sends a <attr_val> (case specific in your schema) as a
   `basic_message`.
1. The *Seller* role sends the same sessionID directly to the *buyer* role. The
   communication channel can be their existing DIDComm connection or something
   else, but the *buyer* needs to know how to react to that line if it's a
   `basic_message`.
1. The *Buyer* role or app implementation for the role sends the received
   `sessionID` to the chatbot, i.e., joins the same session.
1. The *Buyer* role sends the `rcvr` word to the chatbot to make explicit role
   selection. (We could leave this out in some version of FSM implementation
   and rely only on the order of the messages, but this allows us to understand
   better and keep things open for future extensions.)
1. The *Rcvr FSM* instance has now got the actual credential holder
   (Buyer/Receiver) and it sends a `receiver_arriwed` string to the Backend FSM.
1. The *Backend FSM* sends a `rcvr_arrriwed` to the *Issuer FSM* as a
   `basic_message`.
1. Now the *Issuer FSM* loops thru all previously received (from *Seller*)
   attribute values and sends them to the *Backend FSM*.
1. The *Backend FSM* sends the attribute values to the *Rcvr FSM* as a
   `basic_messages`
1. **Optional for the future**: if there would be more attributes than one, this
   would be the place to send the information that all attributes are sent to
   the Backend FSM. Another way to implement these state machines would be to add
   information to both *Issuing* and *Receiving FSMs* how many attributes there
   are, and receiving states would be declared to rely on that knowledge.
1. **Optional**: see the previous step. The *Backend* FSM works as a forwarder
   for all of the cases where the issuing and the receiving FSM instances need
   to communicate with each other through the chatbot service. 
1. Finally the *RcvrFMS* executes **credential issuing protocol**.

