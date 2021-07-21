// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"fmt"
	"time"

	"github.com/textileio/broker-core/broker"
)

type BatchStatus string

const (
	BatchStatusOpen      BatchStatus = "open"
	BatchStatusReady     BatchStatus = "ready"
	BatchStatusExecuting BatchStatus = "executing"
	BatchStatusDone      BatchStatus = "done"
)

func (e *BatchStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BatchStatus(s)
	case string:
		*e = BatchStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BatchStatus: %T", src)
	}
	return nil
}

type Batch struct {
	BatchID   broker.StorageDealID `json:"batchID"`
	Status    BatchStatus          `json:"status"`
	TotalSize int64                `json:"totalSize"`
	CreatedAt time.Time            `json:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt"`
}

type StorageRequest struct {
	OperationID      string                 `json:"operationID"`
	StorageRequestID broker.BrokerRequestID `json:"storageRequestID"`
	DataCid          string                 `json:"dataCid"`
	BatchID          broker.StorageDealID   `json:"batchID"`
	Size             int64                  `json:"size"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
}
