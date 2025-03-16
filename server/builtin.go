package server

import (
	"github.com/ether-edge/ether-edge/chain"
	"github.com/ether-edge/ether-edge/consensus"
	consensusDev "github.com/ether-edge/ether-edge/consensus/dev"
	consensusDummy "github.com/ether-edge/ether-edge/consensus/dummy"
	consensusIBFT "github.com/ether-edge/ether-edge/consensus/ibft"
	consensusPolyBFT "github.com/ether-edge/ether-edge/consensus/polybft"
	"github.com/ether-edge/ether-edge/forkmanager"
	"github.com/ether-edge/ether-edge/secrets"
	"github.com/ether-edge/ether-edge/secrets/awsssm"
	"github.com/ether-edge/ether-edge/secrets/gcpssm"
	"github.com/ether-edge/ether-edge/secrets/hashicorpvault"
	"github.com/ether-edge/ether-edge/secrets/local"
	"github.com/ether-edge/ether-edge/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
