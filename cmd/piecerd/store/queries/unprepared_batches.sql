-- name: CreateUnpreparedBatch :exec
INSERT INTO unprepared_batches(
    storage_deal_id,
    status,
    data_cid
 ) VALUES ($1, $2, $3);

-- name: GetNextPending :one
UPDATE unprepared_batches
SET status = 2, updated_at = CURRENT_TIMESTAMP
WHERE storage_deal_id = (SELECT ub.storage_deal_id FROM unprepared_batches ub
            WHERE ub.ready_at < CURRENT_TIMESTAMP AND
                  ub.status = 1  
                  ORDER BY ub.ready_at asc 
                  FOR UPDATE SKIP LOCKED
                  LIMIT 1)
RETURNING *;

-- name: DeleteUnpreparedBatch :execrows
DELETE FROM unprepared_batches 
WHERE storage_deal_id = $1 AND 
      status = 2;


-- name: MoveToPending :execrows
UPDATE unprepared_batches 
SET status = 1, updated_at = CURRENT_TIMESTAMP, ready_at=$2
WHERE storage_deal_id = $1 AND 
      status = 2;


