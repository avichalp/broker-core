package service_test

import (
	"crypto/rand"
	"path/filepath"
	"testing"

	golog "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	core "github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/bidbot/service"
	"github.com/textileio/broker-core/dshelper"
	"github.com/textileio/broker-core/logging"
	"github.com/textileio/broker-core/marketpeer"
)

const (
	oneDayEpochs = 60 * 24 * 2
)

func init() {
	if err := logging.SetLogLevels(map[string]golog.LogLevel{
		"bidbot/service": golog.LevelDebug,
		"bidbot/store":   golog.LevelDebug,
		"mpeer":          golog.LevelDebug,
	}); err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	dir := t.TempDir()

	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	bidParams := service.BidParams{
		WalletAddrSig:            []byte("bar"),
		AskPrice:                 100000000000,
		VerifiedAskPrice:         100000000000,
		FastRetrieval:            true,
		DealStartWindow:          oneDayEpochs,
		ProposalCidFetchAttempts: 3,
		ProposalDataDirectory:    t.TempDir(),
	}
	auctionFilters := service.AuctionFilters{
		DealDuration: service.MinMaxFilter{
			Min: core.MinDealEpochs,
			Max: core.MaxDealEpochs,
		},
		DealSize: service.MinMaxFilter{
			Min: 56 * 1024,
			Max: 32 * 1000 * 1000 * 1000,
		},
	}

	config := service.Config{
		Peer: marketpeer.Config{
			PrivKey:    priv,
			RepoPath:   dir,
			EnableMDNS: true,
		},
	}

	store, err := dshelper.NewBadgerTxnDatastore(filepath.Join(dir, "bidstore"))
	require.NoError(t, err)
	defer func() { _ = store.Close() }()

	fc := newFilClientMock()

	config.AuctionFilters = auctionFilters

	// Bad DealStartWindow
	config.BidParams = service.BidParams{
		DealStartWindow:          0,
		ProposalCidFetchAttempts: 1,
		ProposalDataDirectory:    t.TempDir(),
	}
	_, err = service.New(config, store, fc)
	require.Error(t, err)

	// Bad ProposalCidFetchAttempts
	config.BidParams = service.BidParams{
		DealStartWindow:          oneDayEpochs,
		ProposalCidFetchAttempts: 0,
		ProposalDataDirectory:    t.TempDir(),
	}
	_, err = service.New(config, store, fc)
	require.Error(t, err)

	// Bad ProposalDataDirectory
	config.BidParams = service.BidParams{
		DealStartWindow:          oneDayEpochs,
		ProposalCidFetchAttempts: 1,
		ProposalDataDirectory:    "",
	}
	_, err = service.New(config, store, fc)
	require.Error(t, err)

	config.BidParams = bidParams

	// Bad auction MinMaxFilter
	config.AuctionFilters = service.AuctionFilters{
		DealDuration: service.MinMaxFilter{
			Min: 10, // min greater than max
			Max: 0,
		},
		DealSize: service.MinMaxFilter{
			Min: 10,
			Max: 20,
		},
	}
	_, err = service.New(config, store, fc)
	require.Error(t, err)

	config.AuctionFilters = auctionFilters

	// Good config
	s, err := service.New(config, store, fc)
	require.NoError(t, err)
	err = s.Subscribe(false)
	require.NoError(t, err)
	require.NoError(t, s.Close())

	// Ensure verify bidder can lead to error
	fc2 := &fcMock{}
	fc2.On(
		"VerifyBidder",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(false, nil)
	_, err = service.New(config, store, fc2)
	require.Error(t, err)
}

func newFilClientMock() *fcMock {
	cm := &fcMock{}
	cm.On(
		"VerifyBidder",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(true, nil)
	cm.On("GetChainHeight").Return(uint64(0), nil)
	return cm
}

type fcMock struct {
	mock.Mock
}

func (fc *fcMock) Close() error {
	args := fc.Called()
	return args.Error(0)
}

func (fc *fcMock) VerifyBidder(bidderSig []byte, bidderID peer.ID, minerAddr string) (bool, error) {
	args := fc.Called(bidderSig, bidderID, minerAddr)
	return args.Bool(0), args.Error(1)
}

func (fc *fcMock) GetChainHeight() (uint64, error) {
	args := fc.Called()
	return args.Get(0).(uint64), args.Error(1)
}
