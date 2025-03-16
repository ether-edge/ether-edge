package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ether-edge/ether-edge/command/backup"
	"github.com/ether-edge/ether-edge/command/bridge"
	"github.com/ether-edge/ether-edge/command/genesis"
	"github.com/ether-edge/ether-edge/command/helper"
	"github.com/ether-edge/ether-edge/command/ibft"
	"github.com/ether-edge/ether-edge/command/license"
	"github.com/ether-edge/ether-edge/command/monitor"
	"github.com/ether-edge/ether-edge/command/peers"
	"github.com/ether-edge/ether-edge/command/polybft"
	"github.com/ether-edge/ether-edge/command/polybftsecrets"
	"github.com/ether-edge/ether-edge/command/regenesis"
	"github.com/ether-edge/ether-edge/command/rootchain"
	"github.com/ether-edge/ether-edge/command/secrets"
	"github.com/ether-edge/ether-edge/command/server"
	"github.com/ether-edge/ether-edge/command/status"
	"github.com/ether-edge/ether-edge/command/txpool"
	"github.com/ether-edge/ether-edge/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
