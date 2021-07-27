// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addStorageRequestInBatchStmt, err = db.PrepareContext(ctx, addStorageRequestInBatch); err != nil {
		return nil, fmt.Errorf("error preparing query AddStorageRequestInBatch: %w", err)
	}
	if q.createOpenBatchStmt, err = db.PrepareContext(ctx, createOpenBatch); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOpenBatch: %w", err)
	}
	if q.doneBatchStatsStmt, err = db.PrepareContext(ctx, doneBatchStats); err != nil {
		return nil, fmt.Errorf("error preparing query DoneBatchStats: %w", err)
	}
	if q.findOpenBatchWithSpaceStmt, err = db.PrepareContext(ctx, findOpenBatchWithSpace); err != nil {
		return nil, fmt.Errorf("error preparing query FindOpenBatchWithSpace: %w", err)
	}
	if q.getNextReadyBatchStmt, err = db.PrepareContext(ctx, getNextReadyBatch); err != nil {
		return nil, fmt.Errorf("error preparing query GetNextReadyBatch: %w", err)
	}
	if q.getStorageRequestsFromBatchStmt, err = db.PrepareContext(ctx, getStorageRequestsFromBatch); err != nil {
		return nil, fmt.Errorf("error preparing query GetStorageRequestsFromBatch: %w", err)
	}
	if q.moveBatchToStatusStmt, err = db.PrepareContext(ctx, moveBatchToStatus); err != nil {
		return nil, fmt.Errorf("error preparing query MoveBatchToStatus: %w", err)
	}
	if q.openBatchStatsStmt, err = db.PrepareContext(ctx, openBatchStats); err != nil {
		return nil, fmt.Errorf("error preparing query OpenBatchStats: %w", err)
	}
	if q.updateBatchSizeStmt, err = db.PrepareContext(ctx, updateBatchSize); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateBatchSize: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addStorageRequestInBatchStmt != nil {
		if cerr := q.addStorageRequestInBatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addStorageRequestInBatchStmt: %w", cerr)
		}
	}
	if q.createOpenBatchStmt != nil {
		if cerr := q.createOpenBatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOpenBatchStmt: %w", cerr)
		}
	}
	if q.doneBatchStatsStmt != nil {
		if cerr := q.doneBatchStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing doneBatchStatsStmt: %w", cerr)
		}
	}
	if q.findOpenBatchWithSpaceStmt != nil {
		if cerr := q.findOpenBatchWithSpaceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findOpenBatchWithSpaceStmt: %w", cerr)
		}
	}
	if q.getNextReadyBatchStmt != nil {
		if cerr := q.getNextReadyBatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNextReadyBatchStmt: %w", cerr)
		}
	}
	if q.getStorageRequestsFromBatchStmt != nil {
		if cerr := q.getStorageRequestsFromBatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStorageRequestsFromBatchStmt: %w", cerr)
		}
	}
	if q.moveBatchToStatusStmt != nil {
		if cerr := q.moveBatchToStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing moveBatchToStatusStmt: %w", cerr)
		}
	}
	if q.openBatchStatsStmt != nil {
		if cerr := q.openBatchStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing openBatchStatsStmt: %w", cerr)
		}
	}
	if q.updateBatchSizeStmt != nil {
		if cerr := q.updateBatchSizeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateBatchSizeStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	addStorageRequestInBatchStmt    *sql.Stmt
	createOpenBatchStmt             *sql.Stmt
	doneBatchStatsStmt              *sql.Stmt
	findOpenBatchWithSpaceStmt      *sql.Stmt
	getNextReadyBatchStmt           *sql.Stmt
	getStorageRequestsFromBatchStmt *sql.Stmt
	moveBatchToStatusStmt           *sql.Stmt
	openBatchStatsStmt              *sql.Stmt
	updateBatchSizeStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		addStorageRequestInBatchStmt:    q.addStorageRequestInBatchStmt,
		createOpenBatchStmt:             q.createOpenBatchStmt,
		doneBatchStatsStmt:              q.doneBatchStatsStmt,
		findOpenBatchWithSpaceStmt:      q.findOpenBatchWithSpaceStmt,
		getNextReadyBatchStmt:           q.getNextReadyBatchStmt,
		getStorageRequestsFromBatchStmt: q.getStorageRequestsFromBatchStmt,
		moveBatchToStatusStmt:           q.moveBatchToStatusStmt,
		openBatchStatsStmt:              q.openBatchStatsStmt,
		updateBatchSizeStmt:             q.updateBatchSizeStmt,
	}
}