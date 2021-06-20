package filclient

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/broker-core/finalizer"
)

var requestTimeout = time.Second * 10

// FilClient provides functionalities to verify bidders.
type FilClient struct {
	fullNode v0api.FullNode
	fakeMode bool

	ctx       context.Context
	finalizer *finalizer.Finalizer
}

// New returns a new FilClient.
func New(lotusGatewayURL string, fakeMode bool) (*FilClient, error) {
	fin := finalizer.NewFinalizer()
	ctx, cancel := context.WithCancel(context.Background())
	fin.Add(finalizer.NewContextCloser(cancel))

	var fn v0api.FullNodeStruct
	fncloser, err := jsonrpc.NewClient(ctx, lotusGatewayURL, "Filecoin", &fn.Internal, http.Header{})
	if err != nil {
		return nil, fmt.Errorf("creating fullnode json rpc client: %v", err)
	}
	fin.AddFn(fncloser)

	return &FilClient{
		fullNode:  &fn,
		fakeMode:  fakeMode,
		ctx:       ctx,
		finalizer: fin,
	}, nil
}

// Close the client.
func (fc *FilClient) Close() error {
	return fc.finalizer.Cleanup(nil)
}

// VerifyBidder ensures that the wallet address authorized the use of bidder peer.ID to make bids.
// Miner's authorize a bidding peer.ID by signing it with a wallet address private key.
func (fc *FilClient) VerifyBidder(bidderSig []byte, bidderID peer.ID, minerAddrStr string) (bool, error) {
	if fc.fakeMode {
		return true, nil
	}
	var sig crypto.Signature
	err := sig.UnmarshalBinary(bidderSig)
	if err != nil {
		return false, fmt.Errorf("unmarshaling signature: %v", err)
	}

	minerAddr, err := address.NewFromString(minerAddrStr)
	if err != nil {
		return false, fmt.Errorf("parsing miner address: %v", err)
	}
	ctx, cancel := context.WithTimeout(fc.ctx, requestTimeout)
	defer cancel()
	mi, err := fc.fullNode.StateMinerInfo(ctx, minerAddr, types.EmptyTSK)
	if err != nil {
		return false, fmt.Errorf("calling full node state miner info: %v", err)
	}
	ownerWalletAddr, err := fc.fullNode.StateAccountKey(ctx, mi.Owner, types.EmptyTSK)
	if err != nil {
		return false, fmt.Errorf("calling full node state account key: %v", err)
	}

	ctx, cancel = context.WithTimeout(fc.ctx, requestTimeout)
	defer cancel()
	ok, err := fc.fullNode.WalletVerify(ctx, ownerWalletAddr, []byte(bidderID), &sig)
	if err != nil {
		return false, fmt.Errorf("calling full node wallet verify: %v", err)
	}
	return ok, nil
}

// GetChainHeight returns the current chain height in epochs.
func (fc *FilClient) GetChainHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(fc.ctx, requestTimeout)
	defer cancel()
	ts, err := fc.fullNode.ChainHead(ctx)
	if err != nil {
		return 0, fmt.Errorf("calling full node chain head: %v", err)
	}
	return uint64(ts.Height()), nil
}
