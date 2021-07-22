-- name: CreateAuctionDeal :exec
INSERT INTO auction_deals(
id,
auction_data_id,
miner_id,
price_per_gib_per_epoch,
start_epoch,
verified,
fast_retrieval,
auction_id,
bid_id,
status,
executing,
error_cause,
retries,
proposal_cid,
deal_id,
deal_expiration,
deal_market_status,
ready_at
    ) VALUES(
      $1,
      $2,
      $3,
      $4,
      $5,
      $6,
      $7,
      $8,
      $9,
      $10,
      $11,
      $12,
      $13,
      $14,
      $15,
      $16,
      $17,
      $18
      );

-- name: NextPendingAuctionDeal :one
UPDATE auction_deals
SET executing = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = (SELECT id FROM auction_deals
    WHERE auction_deals.ready_at < CURRENT_TIMESTAMP AND
          auction_deals.status=$1 AND
          NOT auction_deals.executing
    ORDER BY auction_deals.ready_at asc
    FOR UPDATE SKIP LOCKED
    LIMIT 1)
RETURNING *;

-- name: UpdateAuctionDeal :execrows
UPDATE auction_deals
SET 
    auction_data_id = @auction_data_id,
    miner_id = @miner_id,
    price_per_gib_per_epoch = @price_per_gib_per_epoch,
    start_epoch = @start_epoch,
    verified = @verified,
    fast_retrieval = @fast_retrieval,
    auction_id = @auction_id,
    bid_id = @bid_id,
    status = @status,
    executing = @executing,
    error_cause = @error_cause,
    retries = @retries,
    proposal_cid = @proposal_cid,
    deal_id = @deal_id,
    deal_expiration = @deal_expiration,
    deal_market_status = @deal_market_status,
    ready_at = @ready_at,
    updated_at = CURRENT_TIMESTAMP
    WHERE id = @id;

-- name: GetAuctionDeal :one
SELECT * FROM auction_deals WHERE id = $1;

-- name: GetAuctionDealIDs :many
SELECT id FROM auction_deals WHERE auction_data_id = $1;

-- name: GetAuctionDealsByStatus :many
SELECT * FROM auction_deals WHERE status = $1;

-- name: RemoveAuctionDeal :exec
DELETE FROM auction_deals WHERE id = $1;