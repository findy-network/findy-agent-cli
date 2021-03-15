package main

import (
	"github.com/findy-network/findy-agent-cli/cmd"
	_ "github.com/findy-network/findy-agent-cli/cmd/authn"
	_ "github.com/findy-network/findy-agent-cli/cmd/connection/bot"
	_ "github.com/findy-network/findy-agent-cli/cmd/ops"
)

func main() {
	cmd.Execute()
}
