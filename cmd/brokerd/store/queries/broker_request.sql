-- name: CreateBrokerRequest :exec
INSERT INTO broker_requests(
    id,
    data_cid,
    status
    ) VALUES ($1, $2, $3);

-- name: GetBrokerRequest :one
SELECT * FROM broker_requests
WHERE id = $1;


-- name: GetBrokerRequests :many
SELECT id FROM broker_requests
WHERE storage_deal_id = $1;

-- name: GetBrokerRequestsFull :many
SELECT * FROM broker_requests
WHERE storage_deal_id = $1;

-- name: UpdateBrokerRequests :exec
UPDATE broker_requests
SET status = $2,
    storage_deal_id = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = any ($1::TEXT[]);

-- name: UpdateBrokerRequestsStatus :exec
UPDATE broker_requests
SET status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE storage_deal_id = $1;

-- name: RebatchBrokerRequests :exec
UPDATE broker_requests
SET rebatch_count = rebatch_count + 1,
    error_cause = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE storage_deal_id = $1;
