package blockscanner

import "osmosis_bridge/bridge/bifrost/thorclient/types"

type DummyFetcher struct {
	Tx  types.TxIn
	Err error
}

func NewDummyFetcher(tx types.TxIn, err error) DummyFetcher {
	return DummyFetcher{
		Tx:  tx,
		Err: err,
	}
}

func (d DummyFetcher) FetchMemPool(height int64) (types.TxIn, error) {
	return d.Tx, d.Err
}

func (d DummyFetcher) FetchTxs(height, _ int64) (types.TxIn, error) {
	return d.Tx, d.Err
}

func (d DummyFetcher) GetHeight() (int64, error) {
	return 0, nil
}
