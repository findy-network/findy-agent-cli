package main

import (
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	_ "github.com/findy-network/findy-agent-cli/cmd/agency"
	_ "github.com/findy-network/findy-agent-cli/cmd/agent"
	_ "github.com/findy-network/findy-agent-cli/cmd/authn"
	_ "github.com/findy-network/findy-agent-cli/cmd/bot"
	_ "github.com/findy-network/findy-agent-cli/cmd/connection"
	"github.com/lainio/err2"
)

func main() {
	err2.SetPanicTracer(os.Stderr)
	cmd.Execute()
}
