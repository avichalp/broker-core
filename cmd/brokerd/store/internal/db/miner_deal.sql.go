// Code generated by sqlc. DO NOT EDIT.
// source: miner_deal.sql

package db

import (
	"context"

	"github.com/textileio/bidbot/lib/auction"
	"github.com/textileio/broker-core/broker"
)

const createMinerDeal = `-- name: CreateMinerDeal :exec
INSERT INTO miner_deals(
    storage_deal_id,
    auction_id,
    bid_id,
    miner_addr,
    deal_id,
    deal_expiration,
    error_cause
    ) VALUES(
      $1,
      $2,
      $3,
      $4,
      $5,
      $6,
      $7
      )
`

type CreateMinerDealParams struct {
	StorageDealID  broker.StorageDealID `json:"storageDealID"`
	AuctionID      auction.AuctionID    `json:"auctionID"`
	BidID          auction.BidID        `json:"bidID"`
	MinerAddr      string               `json:"minerAddr"`
	DealID         int64                `json:"dealID"`
	DealExpiration uint64               `json:"dealExpiration"`
	ErrorCause     string               `json:"errorCause"`
}

func (q *Queries) CreateMinerDeal(ctx context.Context, arg CreateMinerDealParams) error {
	_, err := q.exec(ctx, q.createMinerDealStmt, createMinerDeal,
		arg.StorageDealID,
		arg.AuctionID,
		arg.BidID,
		arg.MinerAddr,
		arg.DealID,
		arg.DealExpiration,
		arg.ErrorCause,
	)
	return err
}

const getMinerDeals = `-- name: GetMinerDeals :many
SELECT storage_deal_id, auction_id, bid_id, miner_addr, deal_id, deal_expiration, error_cause, created_at, updated_at FROM miner_deals WHERE storage_deal_id = $1
`

func (q *Queries) GetMinerDeals(ctx context.Context, storageDealID broker.StorageDealID) ([]MinerDeal, error) {
	rows, err := q.query(ctx, q.getMinerDealsStmt, getMinerDeals, storageDealID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MinerDeal
	for rows.Next() {
		var i MinerDeal
		if err := rows.Scan(
			&i.StorageDealID,
			&i.AuctionID,
			&i.BidID,
			&i.MinerAddr,
			&i.DealID,
			&i.DealExpiration,
			&i.ErrorCause,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMinerDeals = `-- name: UpdateMinerDeals :many
UPDATE miner_deals
SET deal_id = $3,
    deal_expiration = $4,
    error_cause = $5
WHERE storage_deal_id = $1 AND miner_addr = $2
RETURNING deal_id
`

type UpdateMinerDealsParams struct {
	StorageDealID  broker.StorageDealID `json:"storageDealID"`
	MinerAddr      string               `json:"minerAddr"`
	DealID         int64                `json:"dealID"`
	DealExpiration uint64               `json:"dealExpiration"`
	ErrorCause     string               `json:"errorCause"`
}

func (q *Queries) UpdateMinerDeals(ctx context.Context, arg UpdateMinerDealsParams) ([]int64, error) {
	rows, err := q.query(ctx, q.updateMinerDealsStmt, updateMinerDeals,
		arg.StorageDealID,
		arg.MinerAddr,
		arg.DealID,
		arg.DealExpiration,
		arg.ErrorCause,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var deal_id int64
		if err := rows.Scan(&deal_id); err != nil {
			return nil, err
		}
		items = append(items, deal_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
