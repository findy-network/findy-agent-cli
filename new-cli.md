## New CLI

We will communicate by;
- Agency: set impl, set devops, create schema, how about pump?
- Cloud Agent: listen, read, ping, connecti, invitation, protocols, ...

note! if we would want to onboard we should do that thru the Vault service which
will give us JWT and does the registration. For that reason, that willi be done
thru 'authn' cmd family.

### We will use common.Client

We will use grpc.Client wrapper and bot. As well, authn service, from our wallet
allocation service.

### TODO
- [x] move authn commands to root level
- [x] check offer.go and proof.go: they are old!! **remove**
- [x] think about name for agency cmds with new gRPC API
- [x] move flags to new cmd tree like server addr and port to root cmd
