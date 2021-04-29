package nearclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
	logging "github.com/ipfs/go-log/v2"
	"github.com/textileio/broker-core/cmd/neard/nearclient/account"
	itypes "github.com/textileio/broker-core/cmd/neard/nearclient/internal/types"
	"github.com/textileio/broker-core/cmd/neard/nearclient/types"
	"github.com/textileio/broker-core/cmd/neard/nearclient/util"
)

var (
	log = logging.Logger("nearclient")
)

// CallFunctionResponse holds information about the result of a function call.
type CallFunctionResponse struct {
	Result      []byte   `json:"result"`
	Logs        []string `json:"logs"`
	BlockHeight int      `json:"block_height"`
	BlockHash   string   `json:"block_hash"`
}

// Change holds information about a state change of a key-value pair.
type Change struct {
	AccountID   string `json:"account_id"`
	KeyBase64   string `json:"key_base64"`
	ValueBase64 string `json:"value_base64"`
}

// Cause holds information about the cause of a state change.
type Cause struct {
	Type        string `json:"type"`
	ReceiptHash string `json:"receipt_hash"`
}

// ChangeData holds information about a state change.
type ChangeData struct {
	Cause  Cause  `json:"cause"`
	Type   string `json:"type"`
	Change Change `json:"change"`
}

// DataChangesResponse holds information about all data changes in a block.
type DataChangesResponse struct {
	BlockHash string       `json:"block_hash"`
	Changes   []ChangeData `json:"changes"`
}

// Client communicates with the NEAR API.
type Client struct {
	config *types.Config
}

// NewClient creates a new Client.
func NewClient(config *types.Config) (*Client, error) {
	return &Client{
		config: config,
	}, nil
}

// Account provides an API for the provided account ID.
func (c *Client) Account(accountID string) *account.Account {
	return account.NewAccount(c.config, accountID)
}

// CallFunctionOption controls the behavior when calling CallFunction.
type CallFunctionOption func(*itypes.QueryRequest) error

// CallFunctionWithFinality specifies the finality to be used when calling the function.
func CallFunctionWithFinality(finalaity string) CallFunctionOption {
	return func(qr *itypes.QueryRequest) error {
		qr.Finality = finalaity
		return nil
	}
}

// CallFunctionWithBlockHeight specifies the block height to call the function for.
func CallFunctionWithBlockHeight(blockHeight int) CallFunctionOption {
	return func(qr *itypes.QueryRequest) error {
		qr.BlockID = blockHeight
		return nil
	}
}

// CallFunctionWithBlockHash specifies the block hash to call the function for.
func CallFunctionWithBlockHash(blockHash string) CallFunctionOption {
	return func(qr *itypes.QueryRequest) error {
		qr.BlockID = blockHash
		return nil
	}
}

// CallFunctionWithArgs specified the args to call the function with.
// Should be a JSON encodable object.
func CallFunctionWithArgs(args interface{}) CallFunctionOption {
	return func(qr *itypes.QueryRequest) error {
		if args == nil {
			args = make(map[string]interface{})
		}
		bytes, err := json.Marshal(args)
		if err != nil {
			return err
		}
		qr.ArgsBase64 = base64.StdEncoding.EncodeToString(bytes)
		return nil
	}
}

// CallFunction calls a function on a contract.
func (c *Client) CallFunction(
	ctx context.Context,
	accountID string,
	methodName string,
	opts ...CallFunctionOption,
) (*CallFunctionResponse, error) {
	req := &itypes.QueryRequest{
		RequestType: "call_function",
		AccountID:   accountID,
		MethodName:  methodName,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.BlockID == nil && req.Finality == "" {
		return nil, fmt.Errorf(
			"you must provide ViewAccountWithBlockHeight, ViewAccountWithBlockHash or ViewAccountWithFinality",
		)
	}
	if req.BlockID != nil && req.Finality != "" {
		return nil, fmt.Errorf(
			"you must provide one of ViewAccountWithBlockHeight, ViewAccountWithBlockHash or ViewAccountWithFinality",
		)
	}
	var res CallFunctionResponse
	if err := c.config.RPCClient.CallContext(ctx, &res, "query", rpc.NewNamedParams(req)); err != nil {
		return nil, fmt.Errorf("calling query rpc: %v", util.MapRPCError(err))
	}
	return &res, nil
}

// DataChangesOption controls behavior when calling DataChanges.
type DataChangesOption func(*itypes.ChangesRequest)

// DataChangesWithPrefix sets the data key prefix to query for.
func DataChangesWithPrefix(prefix string) DataChangesOption {
	return func(cr *itypes.ChangesRequest) {
		cr.KeyPrefixBase64 = base64.StdEncoding.EncodeToString([]byte(prefix))
	}
}

// DataChangesWithFinality specifies the finality to be used when querying data changes.
func DataChangesWithFinality(finalaity string) DataChangesOption {
	return func(qr *itypes.ChangesRequest) {
		qr.Finality = finalaity
	}
}

// DataChangesWithBlockHeight specifies the block id to query data changes for.
func DataChangesWithBlockHeight(blockHeight int) DataChangesOption {
	return func(qr *itypes.ChangesRequest) {
		qr.BlockID = blockHeight
	}
}

// DataChangesWithBlockHash specifies the block id to query data changes for.
func DataChangesWithBlockHash(blockHash string) DataChangesOption {
	return func(qr *itypes.ChangesRequest) {
		qr.BlockID = blockHash
	}
}

// DataChanges queries changes to contract data changes.
func (c *Client) DataChanges(
	ctx context.Context,
	accountIDs []string,
	opts ...DataChangesOption,
) (*DataChangesResponse, error) {
	req := &itypes.ChangesRequest{
		ChangesType: "data_changes",
		AccountIDs:  accountIDs,
	}
	for _, opt := range opts {
		opt(req)
	}
	if req.BlockID == nil && req.Finality == "" {
		return nil, fmt.Errorf(
			"you must provide DataChangesWithBlockHeight, DataChangesWithBlockHash or DataChangesWithFinality",
		)
	}
	if req.BlockID != nil && req.Finality != "" {
		return nil, fmt.Errorf(
			"you must provide one of DataChangesWithBlockHeight, DataChangesWithBlockHash or DataChangesWithFinality",
		)
	}
	var res DataChangesResponse
	if err := c.config.RPCClient.CallContext(ctx, &res, "EXPERIMENTAL_changes", rpc.NewNamedParams(req)); err != nil {
		return nil, fmt.Errorf("calling changes rpc: %v", util.MapRPCError(err))
	}
	return &res, nil
}
