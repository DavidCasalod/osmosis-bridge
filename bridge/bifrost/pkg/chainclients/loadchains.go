package chainclients

import (
	"osmosis_bridge/bridge/common"
	"time"

	"gitlab.com/thorchain/thornode/config"

	"github.com/rs/zerolog/log"
	"gitlab.com/thorchain/tss/go-tss/tss"

	"gitlab.com/thorchain/thornode/bifrost/metrics"
	"gitlab.com/thorchain/thornode/bifrost/pkg/chainclients/bitcoin"
	"gitlab.com/thorchain/thornode/bifrost/pkg/chainclients/gaia"
	"gitlab.com/thorchain/thornode/bifrost/pkg/chainclients/shared/types"
	"gitlab.com/thorchain/thornode/bifrost/pubkeymanager"
	"gitlab.com/thorchain/thornode/bifrost/thorclient"
)

// ChainClient exports the shared type.
type ChainClient = types.ChainClient

// LoadChains returns chain clients from chain configuration
func LoadChains(thorKeys *thorclient.Keys,
	cfg map[common.Chain]config.BifrostChainConfiguration,
	server *tss.TssServer,
	thorchainBridge thorclient.ThorchainBridge,
	m *metrics.Metrics,
	pubKeyValidator pubkeymanager.PubKeyValidator,
	poolMgr thorclient.PoolManager,
) (chains map[common.Chain]ChainClient, restart chan struct{}) {
	logger := log.Logger.With().Str("module", "bifrost").Logger()

	chains = make(map[common.Chain]ChainClient)
	restart = make(chan struct{})
	failedChains := []common.Chain{}

	loadChain := func(chain config.BifrostChainConfiguration) (ChainClient, error) {

		switch common.Chain(chain.ChainID) {

		case common.OSMOSISChain:
			return gaia.NewCosmosClient(thorKeys, chain, server, thorchainBridge, m)
		case common.GAIAChain:
			return gaia.NewCosmosClient(thorKeys, chain, server, thorchainBridge, m)
		case common.BTCChain:
			return bitcoin.NewClient(thorKeys, chain, server, thorchainBridge, m)

		default:
			log.Fatal().Msgf("chain %s is not supported", chain.ChainID)
			return nil, nil
		}
	}

	for _, chain := range cfg {
		if chain.Disabled {
			logger.Info().Msgf("%s chain is disabled by configure", chain.ChainID)
			continue
		}

		client, err := loadChain(chain)
		if err != nil {
			logger.Error().Err(err).Stringer("chain", chain.ChainID).Msg("failed to load chain")
			failedChains = append(failedChains, common.Chain(chain.ChainID))
			continue
		}

		// trunk-ignore-all(golangci-lint/forcetypeassert)
		switch common.Chain(chain.ChainID) {
		case common.BTCChain:
			pubKeyValidator.RegisterCallback(client.(*bitcoin.Client).RegisterPublicKey)

		}
		chains[common.Chain(chain.ChainID)] = client
	}

	// watch failed chains minutely and restart bifrost if any succeed init
	if len(failedChains) > 0 {
		go func() {
			tick := time.NewTicker(time.Minute)
			for range tick.C {
				for _, chain := range failedChains {
					ccfg := cfg[chain]
					ccfg.BlockScanner.DBPath = "" // in-memory db

					_, err := loadChain(ccfg)
					if err == nil {
						logger.Info().Stringer("chain", chain).Msg("chain loaded, restarting bifrost")
						close(restart)
						return
					} else {
						logger.Error().Err(err).Stringer("chain", chain).Msg("failed to load chain")
					}
				}
			}
		}()
	}

	return chains, restart
}
