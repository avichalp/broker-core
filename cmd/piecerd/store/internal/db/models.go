// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"fmt"
	"time"

	"github.com/textileio/broker-core/broker"
)

type UnpreparedBatchStatus string

const (
	UnpreparedBatchStatusPending   UnpreparedBatchStatus = "pending"
	UnpreparedBatchStatusExecuting UnpreparedBatchStatus = "executing"
	UnpreparedBatchStatusDone      UnpreparedBatchStatus = "done"
)

func (e *UnpreparedBatchStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnpreparedBatchStatus(s)
	case string:
		*e = UnpreparedBatchStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UnpreparedBatchStatus: %T", src)
	}
	return nil
}

type UnpreparedBatch struct {
	StorageDealID broker.StorageDealID  `json:"storageDealID"`
	Status        UnpreparedBatchStatus `json:"status"`
	DataCid       string                `json:"dataCid"`
	ReadyAt       time.Time             `json:"readyAt"`
	CreatedAt     time.Time             `json:"createdAt"`
	UpdatedAt     time.Time             `json:"updatedAt"`
}