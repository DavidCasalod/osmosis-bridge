package osmosis

import (
	"crypto/x509"
	"fmt"
	"math/big"
	"os"

	"osmosis_bridge/bridge/common"

	thorTypes "gitlab.com/thorchain/thornode/bifrost/thorclient/types"

	stypes "osmosis_bridge/bridge/bifrost/thorclient/types"

	"github.com/cosmos/cosmos-sdk/client"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	thorCommon "gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// buildUnsigned takes a MsgSend and other parameters and returns a txBuilder
// It can be used to simulateTx or as the input to signMsg before BraodcastTx
func buildUnsigned(
	txConfig client.TxConfig,
	msg *btypes.MsgSend,
	pubkey common.PubKey,
	memo string,
	fee ctypes.Coins,
	account uint64,
	sequence uint64,
) (client.TxBuilder, error) {
	cpk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, pubkey.String())
	if err != nil {
		return nil, fmt.Errorf("unable to GetPubKeyFromBech32 from cosmos: %w", err)
	}
	txBuilder := txConfig.NewTxBuilder()

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return nil, fmt.Errorf("unable to SetMsgs on txBuilder: %w", err)
	}

	txBuilder.SetMemo(memo)
	txBuilder.SetFeeAmount(fee)
	txBuilder.SetGasLimit(GasLimit)

	sigData := &signingtypes.SingleSignatureData{
		SignMode: signingtypes.SignMode_SIGN_MODE_DIRECT,
	}
	sig := signingtypes.SignatureV2{
		PubKey:   cpk,
		Data:     sigData,
		Sequence: sequence,
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		return nil, fmt.Errorf("unable to initial SetSignatures on txBuilder: %w", err)
	}

	return txBuilder, nil
}

func fromCosmosToThorchain(c cosmos.Coin) (common.Coin, error) {
	cosmosAsset, exists := GetAssetByCosmosDenom(c.Denom)
	if !exists {
		return common.NoCoin, fmt.Errorf("asset does not exist / not whitelisted by client")
	}

	thorAsset, err := common.NewAsset(fmt.Sprintf("%s.%s", common.OSMOSISChain.String(), cosmosAsset.THORChainSymbol))
	if err != nil {
		return common.NoCoin, fmt.Errorf("invalid thorchain asset: %w", err)
	}

	decimals := cosmosAsset.CosmosDecimals
	amount := c.Amount.BigInt()
	var exp big.Int
	// Decimals are more than native THORChain, so divide...
	if decimals > common.THORChainDecimals {
		decimalDiff := int64(decimals - common.THORChainDecimals)
		amount.Quo(amount, exp.Exp(big.NewInt(10), big.NewInt(decimalDiff), nil))
	} else if decimals < common.THORChainDecimals {
		// Decimals are less than native THORChain, so multiply...
		decimalDiff := int64(common.THORChainDecimals - decimals)
		amount.Mul(amount, exp.Exp(big.NewInt(10), big.NewInt(decimalDiff), nil))
	}
	return common.Coin{
		Asset:    thorAsset,
		Amount:   ctypes.NewUintFromBigInt(amount),
		Decimals: int64(decimals),
	}, nil
}

func fromThorchainToCosmos(coin common.Coin) (cosmos.Coin, error) {
	asset, exists := GetAssetByThorchainSymbol(coin.Asset.Symbol.String())
	if !exists {
		return cosmos.Coin{}, fmt.Errorf("asset does not exist / not whitelisted by client")
	}

	decimals := asset.CosmosDecimals
	amount := coin.Amount.BigInt()
	var exp big.Int
	if decimals > common.THORChainDecimals {
		// Decimals are more than native THORChain, so multiply...
		decimalDiff := int64(decimals - common.THORChainDecimals)
		amount.Mul(amount, exp.Exp(big.NewInt(10), big.NewInt(decimalDiff), nil))
	} else if decimals < common.THORChainDecimals {
		// Decimals are less than native THORChain, so divide...
		decimalDiff := int64(common.THORChainDecimals - decimals)
		amount.Quo(amount, exp.Exp(big.NewInt(10), big.NewInt(decimalDiff), nil))
	}
	return cosmos.NewCoin(asset.CosmosDenom, ctypes.NewIntFromBigInt(amount)), nil
}

func getGRPCConn(host string, tls bool) (*grpc.ClientConn, error) {
	// load system certificates or proceed with insecure if tls disabled
	var creds credentials.TransportCredentials
	if tls {
		certs, err := x509.SystemCertPool()
		if err != nil {
			return &grpc.ClientConn{}, fmt.Errorf("unable to load system certs: %w", err)
		}
		creds = credentials.NewClientTLSFromCert(certs, "")
	} else {
		creds = insecure.NewCredentials()
	}

	return grpc.Dial(host, grpc.WithTransportCredentials(creds))
}

func unmarshalJSONToPb(filePath string, msg proto.Message) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	u := new(jsonpb.Unmarshaler)
	u.AllowUnknownFields = true
	return u.Unmarshal(jsonFile, msg)
}

// ----------- Convert
func ConvertChains(chains common.Chains) thorCommon.Chains {
	convertedChains := make(thorCommon.Chains, len(chains))
	for i, chain := range chains {
		convertedChains[i] = ConvertChain(chain)
	}
	return convertedChains
}
func ConvertToThorAsset(osmosisAsset common.Asset) thorCommon.Asset {
	return thorCommon.Asset{
		Chain:  thorCommon.Chain(osmosisAsset.Chain),
		Symbol: thorCommon.Symbol(osmosisAsset.Symbol),
		Ticker: thorCommon.Ticker(osmosisAsset.Ticker),
		Synth:  osmosisAsset.Synth,
	}
}
func ConvertToThorCoin(osmosisCoins common.Coin) thorCommon.Coin {
	return thorCommon.Coin{
		Asset:    ConvertToThorAsset(osmosisCoins.Asset),
		Amount:   osmosisCoins.Amount,
		Decimals: osmosisCoins.Decimals,
	}

}

func ConvertToThorCoins(osmosisCoins common.Coins) thorCommon.Coins {
	var thorCoins thorCommon.Coins
	for _, coin := range osmosisCoins {
		thorCoin := thorCommon.Coin{
			Asset:    ConvertToThorAsset(coin.Asset),
			Amount:   coin.Amount,
			Decimals: coin.Decimals,
		}
		thorCoins = append(thorCoins, thorCoin)
	}
	return thorCoins
}
func ConvertToThorAccount(osmoAccount common.Account) thorCommon.Account {
	return thorCommon.Account{
		Sequence:      osmoAccount.Sequence,
		AccountNumber: osmoAccount.AccountNumber,
		Coins:         ConvertToThorCoins(osmoAccount.Coins),
		HasMemoFlag:   osmoAccount.HasMemoFlag,
	}
}

func ConvertToOsmCoins(osmosisCoins thorCommon.Coins) common.Coins {
	var osmoCoins common.Coins
	for _, coin := range osmosisCoins {
		osmoCoin := common.Coin{
			Asset:    ConvertToOsmoAsset(coin.Asset),
			Amount:   coin.Amount,
			Decimals: coin.Decimals,
		}
		osmoCoins = append(osmoCoins, osmoCoin)
	}
	return osmoCoins
}

func ConvertToOsmCoin(osmosisCoins thorCommon.Coin) common.Coin {
	return common.Coin{
		Asset:    ConvertToOsmoAsset(osmosisCoins.Asset),
		Amount:   osmosisCoins.Amount,
		Decimals: osmosisCoins.Decimals,
	}
}

func ConvertToOsmoAsset(thorAsset thorCommon.Asset) common.Asset {
	return common.Asset{
		Chain:  common.Chain(thorAsset.Chain),
		Symbol: common.Symbol(thorAsset.Symbol),
		Ticker: common.Ticker(thorAsset.Ticker),
		Synth:  thorAsset.Synth,
	}
}

func ConvertToThorTxIn(txIn stypes.TxIn) thorTypes.TxIn {
	var thorTxArray []thorTypes.TxInItem
	for _, item := range txIn.TxArray {
		thorTxArray = append(thorTxArray, ConvertTxInItem(item))
	}

	return thorTypes.TxIn{
		Count:                txIn.Count,
		Chain:                txIn.Chain,
		TxArray:              thorTxArray,
		Filtered:             txIn.Filtered,
		MemPool:              txIn.MemPool,
		SentUnFinalised:      txIn.SentUnFinalised,
		Finalised:            txIn.Finalised,
		ConfirmationRequired: txIn.ConfirmationRequired,
	}
}
func ConvertTxInItem(item stypes.TxInItem) thorTypes.TxInItem {
	return thorTypes.TxInItem{
		BlockHeight:           item.BlockHeight,
		Tx:                    item.Tx,
		Memo:                  item.Memo,
		Sender:                item.Sender,
		To:                    item.To,
		Coins:                 item.Coins,               // Assuming common.Coins is the same type for both
		Gas:                   item.Gas,                 // Assuming common.Gas is the same type for both
		ObservedVaultPubKey:   item.ObservedVaultPubKey, // Assuming common.PubKey is the same type for both
		Aggregator:            item.Aggregator,
		AggregatorTarget:      item.AggregatorTarget,
		AggregatorTargetLimit: item.AggregatorTargetLimit, // Assuming cosmos.Uint is the same type for both
	}
}
