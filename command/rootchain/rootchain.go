package rootchain

import (
	"github.com/spf13/cobra"

	"github.com/ether-edge/ether-edge/command/rootchain/deploy"
	"github.com/ether-edge/ether-edge/command/rootchain/fund"
	"github.com/ether-edge/ether-edge/command/rootchain/premine"
	"github.com/ether-edge/ether-edge/command/rootchain/server"
)

// GetCommand creates "rootchain" helper command
func GetCommand() *cobra.Command {
	rootchainCmd := &cobra.Command{
		Use:   "rootchain",
		Short: "Top level rootchain helper command.",
	}

	rootchainCmd.AddCommand(
		// rootchain server
		server.GetCommand(),
		// rootchain deploy
		deploy.GetCommand(),
		// rootchain fund
		fund.GetCommand(),
		// rootchain premine
		premine.GetCommand(),
	)

	return rootchainCmd
}
