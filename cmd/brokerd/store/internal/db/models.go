// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/textileio/bidbot/lib/auction"
	"github.com/textileio/broker-core/broker"
)

type Batch struct {
	ID                 broker.BatchID     `json:"id"`
	Status             broker.BatchStatus `json:"status"`
	RepFactor          int                `json:"repFactor"`
	DealDuration       int                `json:"dealDuration"`
	PayloadCid         string             `json:"payloadCid"`
	PieceCid           string             `json:"pieceCid"`
	PieceSize          uint64             `json:"pieceSize"`
	CarUrl             string             `json:"carUrl"`
	CarIpfsCid         string             `json:"carIpfsCid"`
	CarIpfsAddrs       string             `json:"carIpfsAddrs"`
	DisallowRebatching bool               `json:"disallowRebatching"`
	FilEpochDeadline   uint64             `json:"filEpochDeadline"`
	Error              string             `json:"error"`
	Origin             string             `json:"origin"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	Providers          []string           `json:"providers"`
}

type BatchManifest struct {
	BatchID  string `json:"batchID"`
	Manifest []byte `json:"manifest"`
}

type BatchRemoteWallet struct {
	BatchID    broker.BatchID `json:"batchID"`
	PeerID     string         `json:"peerID"`
	AuthToken  string         `json:"authToken"`
	WalletAddr string         `json:"walletAddr"`
	Multiaddrs []string       `json:"multiaddrs"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

type BatchTag struct {
	BatchID   broker.BatchID `json:"batchID"`
	Key       string         `json:"key"`
	Value     string         `json:"value"`
	CreatedAt time.Time      `json:"createdAt"`
}

type Deal struct {
	BatchID           broker.BatchID `json:"batchID"`
	AuctionID         auction.ID     `json:"auctionID"`
	BidID             auction.BidID  `json:"bidID"`
	StorageProviderID string         `json:"storageProviderID"`
	DealID            int64          `json:"dealID"`
	DealExpiration    uint64         `json:"dealExpiration"`
	ErrorCause        string         `json:"errorCause"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

type Operation struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type StorageRequest struct {
	ID           broker.StorageRequestID     `json:"id"`
	DataCid      string                      `json:"dataCid"`
	BatchID      sql.NullString              `json:"batchID"`
	Status       broker.StorageRequestStatus `json:"status"`
	Origin       string                      `json:"origin"`
	RebatchCount int32                       `json:"rebatchCount"`
	ErrorCause   string                      `json:"errorCause"`
	CreatedAt    time.Time                   `json:"createdAt"`
	UpdatedAt    time.Time                   `json:"updatedAt"`
}

type UnpinJob struct {
	ID        string       `json:"id"`
	Executing sql.NullBool `json:"executing"`
	Cid       string       `json:"cid"`
	Type      int16        `json:"type"`
	ReadyAt   time.Time    `json:"readyAt"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
