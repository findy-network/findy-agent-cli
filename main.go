package main

import (
	"github.com/findy-network/findy-agent-cli/cmd"
	_ "github.com/findy-network/findy-agent-cli/extracmd"
	_ "github.com/findy-network/findy-agent-cli/extracmd/impl"
	_ "github.com/findy-network/findy-agent-cli/extracmd/jwt"
	_ "github.com/findy-network/findy-agent-cli/extracmd/ops"
	_ "github.com/findy-network/findy-agent/grpc"
)

func main() {
	cmd.Execute()
}
