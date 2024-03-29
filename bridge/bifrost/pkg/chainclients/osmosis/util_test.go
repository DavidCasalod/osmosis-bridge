package osmosis

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	. "gopkg.in/check.v1"
)

type UtilTestSuite struct{}

var _ = Suite(&UtilTestSuite{})

func (s *UtilTestSuite) SetUpSuite(c *C) {}

func (s *UtilTestSuite) TestFromCosmosToThorchain(c *C) {
	// 5 OSMO, 6 decimals
	cosmosCoin := cosmos.NewCoin("uosmo", ctypes.NewInt(5000000))
	thorchainCoin, err := fromCosmosToThorchain(cosmosCoin)
	c.Assert(err, IsNil)

	// 5 OSMO, 8 decimals
	expectedThorchainAsset, err := common.NewAsset("OSMOSIS.OSMO")
	c.Assert(err, IsNil)
	expectedThorchainAmount := ctypes.NewUint(500000000)
	c.Check(thorchainCoin.Asset.Equals(ConvertToOsmoAsset(expectedThorchainAsset)), Equals, true)
	c.Check(thorchainCoin.Amount.BigInt().Int64(), Equals, expectedThorchainAmount.BigInt().Int64())
	c.Check(thorchainCoin.Decimals, Equals, int64(6))
}

func (s *UtilTestSuite) TestFromThorchainToCosmos(c *C) {
	// 6 OSMOSIS.OSMO, 8 decimals
	thorchainAsset, err := common.NewAsset("OSMOSIS.OSMO")
	c.Assert(err, IsNil)
	thorchainCoin := common.Coin{
		Asset:    thorchainAsset,
		Amount:   cosmos.NewUint(600000000),
		Decimals: 6,
	}
	cosmosCoin, err := fromThorchainToCosmos(ConvertToOsmCoin(thorchainCoin))
	c.Assert(err, IsNil)

	// 6 uosmo, 6 decimals
	expectedCosmosDenom := "uosmo"
	expectedCosmosAmount := int64(6000000)
	c.Check(cosmosCoin.Denom, Equals, expectedCosmosDenom)
	c.Check(cosmosCoin.Amount.Int64(), Equals, expectedCosmosAmount)
}
