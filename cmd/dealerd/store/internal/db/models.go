// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"fmt"
	"time"

	"github.com/textileio/bidbot/lib/auction"
	"github.com/textileio/broker-core/broker"
)

type Status string

const (
	StatusDealMaking      Status = "deal-making"
	StatusConfirmation    Status = "confirmation"
	StatusReportFinalized Status = "report-finalized"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type AuctionDatum struct {
	ID         string         `json:"id"`
	BatchID    broker.BatchID `json:"batchID"`
	PayloadCid string         `json:"payloadCid"`
	PieceCid   string         `json:"pieceCid"`
	PieceSize  uint64         `json:"pieceSize"`
	Duration   uint64         `json:"duration"`
	CreatedAt  time.Time      `json:"createdAt"`
}

type AuctionDeal struct {
	ID                  string        `json:"id"`
	AuctionDataID       string        `json:"auctionDataID"`
	StorageProviderID   string        `json:"storageProviderID"`
	PricePerGibPerEpoch int64         `json:"pricePerGibPerEpoch"`
	StartEpoch          uint64        `json:"startEpoch"`
	Verified            bool          `json:"verified"`
	FastRetrieval       bool          `json:"fastRetrieval"`
	AuctionID           auction.ID    `json:"auctionID"`
	BidID               auction.BidID `json:"bidID"`
	Status              Status        `json:"status"`
	Executing           bool          `json:"executing"`
	ErrorCause          string        `json:"errorCause"`
	Retries             int           `json:"retries"`
	ProposalCid         string        `json:"proposalCid"`
	DealID              int64         `json:"dealID"`
	DealExpiration      uint64        `json:"dealExpiration"`
	DealMarketStatus    uint64        `json:"dealMarketStatus"`
	ReadyAt             time.Time     `json:"readyAt"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
}

type MarketDealStatus struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RemoteWallet struct {
	AuctionDataID string    `json:"auctionDataID"`
	PeerID        string    `json:"peerID"`
	AuthToken     string    `json:"authToken"`
	WalletAddr    string    `json:"walletAddr"`
	Multiaddrs    []string  `json:"multiaddrs"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
