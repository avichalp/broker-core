// Code generated by sqlc. DO NOT EDIT.
// source: broker_request.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/textileio/broker-core/broker"
)

const createBrokerRequest = `-- name: CreateBrokerRequest :exec
INSERT INTO broker_requests(
    id,
    data_cid,
    status
    ) VALUES ($1, $2, $3)
`

type CreateBrokerRequestParams struct {
	ID      broker.BrokerRequestID     `json:"id"`
	DataCid string                     `json:"dataCid"`
	Status  broker.BrokerRequestStatus `json:"status"`
}

func (q *Queries) CreateBrokerRequest(ctx context.Context, arg CreateBrokerRequestParams) error {
	_, err := q.exec(ctx, q.createBrokerRequestStmt, createBrokerRequest, arg.ID, arg.DataCid, arg.Status)
	return err
}

const getBrokerRequest = `-- name: GetBrokerRequest :one
SELECT id, data_cid, storage_deal_id, status, rebatch_count, error_cause, created_at, updated_at FROM broker_requests
WHERE id = $1
`

func (q *Queries) GetBrokerRequest(ctx context.Context, id broker.BrokerRequestID) (BrokerRequest, error) {
	row := q.queryRow(ctx, q.getBrokerRequestStmt, getBrokerRequest, id)
	var i BrokerRequest
	err := row.Scan(
		&i.ID,
		&i.DataCid,
		&i.StorageDealID,
		&i.Status,
		&i.RebatchCount,
		&i.ErrorCause,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBrokerRequests = `-- name: GetBrokerRequests :many
SELECT id FROM broker_requests
WHERE storage_deal_id = $1
`

func (q *Queries) GetBrokerRequests(ctx context.Context, storageDealID sql.NullString) ([]broker.BrokerRequestID, error) {
	rows, err := q.query(ctx, q.getBrokerRequestsStmt, getBrokerRequests, storageDealID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []broker.BrokerRequestID
	for rows.Next() {
		var id broker.BrokerRequestID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBrokerRequestsFull = `-- name: GetBrokerRequestsFull :many
SELECT id, data_cid, storage_deal_id, status, rebatch_count, error_cause, created_at, updated_at FROM broker_requests
WHERE storage_deal_id = $1
`

func (q *Queries) GetBrokerRequestsFull(ctx context.Context, storageDealID sql.NullString) ([]BrokerRequest, error) {
	rows, err := q.query(ctx, q.getBrokerRequestsFullStmt, getBrokerRequestsFull, storageDealID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BrokerRequest
	for rows.Next() {
		var i BrokerRequest
		if err := rows.Scan(
			&i.ID,
			&i.DataCid,
			&i.StorageDealID,
			&i.Status,
			&i.RebatchCount,
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

const rebatchBrokerRequests = `-- name: RebatchBrokerRequests :exec
UPDATE broker_requests
SET rebatch_count = rebatch_count + 1,
    error_cause = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE storage_deal_id = $1
`

type RebatchBrokerRequestsParams struct {
	StorageDealID sql.NullString `json:"storageDealID"`
	ErrorCause    string         `json:"errorCause"`
}

func (q *Queries) RebatchBrokerRequests(ctx context.Context, arg RebatchBrokerRequestsParams) error {
	_, err := q.exec(ctx, q.rebatchBrokerRequestsStmt, rebatchBrokerRequests, arg.StorageDealID, arg.ErrorCause)
	return err
}

const updateBrokerRequests = `-- name: UpdateBrokerRequests :exec
UPDATE broker_requests
SET status = $2,
    storage_deal_id = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = any ($1::TEXT[])
`

type UpdateBrokerRequestsParams struct {
	Column1       []string                   `json:"column1"`
	Status        broker.BrokerRequestStatus `json:"status"`
	StorageDealID sql.NullString             `json:"storageDealID"`
}

func (q *Queries) UpdateBrokerRequests(ctx context.Context, arg UpdateBrokerRequestsParams) error {
	_, err := q.exec(ctx, q.updateBrokerRequestsStmt, updateBrokerRequests, pq.Array(arg.Column1), arg.Status, arg.StorageDealID)
	return err
}

const updateBrokerRequestsStatus = `-- name: UpdateBrokerRequestsStatus :exec
UPDATE broker_requests
SET status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE storage_deal_id = $1
`

type UpdateBrokerRequestsStatusParams struct {
	StorageDealID sql.NullString             `json:"storageDealID"`
	Status        broker.BrokerRequestStatus `json:"status"`
}

func (q *Queries) UpdateBrokerRequestsStatus(ctx context.Context, arg UpdateBrokerRequestsStatusParams) error {
	_, err := q.exec(ctx, q.updateBrokerRequestsStatusStmt, updateBrokerRequestsStatus, arg.StorageDealID, arg.Status)
	return err
}
