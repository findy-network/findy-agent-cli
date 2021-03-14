package main

import (
	"github.com/findy-network/findy-agent-cli/cmd"
	_ "github.com/findy-network/findy-agent-cli/extracmd/jwt"
	_ "github.com/findy-network/findy-agent-cli/extracmd/jwt/authn"
	_ "github.com/findy-network/findy-agent-cli/extracmd/jwt/bot"
	_ "github.com/findy-network/findy-agent-cli/extracmd/ops"
)

func main() {
	cmd.Execute()
}
