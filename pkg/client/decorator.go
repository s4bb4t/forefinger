package client

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"github.com/s4bb4t/forefinger/pkg/models"
	"math/big"
)

// BlockByNumber returns pointer to allocated and initialized models.Block and call error if not nil.
func (c *Client) BlockByNumber(ctx context.Context, number *big.Int) (*models.Block, error) {
	var block models.Block
	return &block, c.Call(ctx, &block, methods.BlockByNumber, number, true)
}

// Balance returns big.Int eth balance of provided address
func (c *Client) Balance(ctx context.Context, address common.Address, block any) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.Balance, address, block)
}

// BlockByHash returns pointer to allocated and initialized models.Block and call error if not nil.
func (c *Client) BlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error) {
	var b models.Block
	return &b, c.Call(ctx, &b, methods.BlockByHash, hash, true)
}

// TxByHash returns pointer to allocated and initialized models.Transaction and call error if not nil.
func (c *Client) TxByHash(ctx context.Context, hash common.Hash) (*models.Transaction, error) {
	var tx models.Transaction
	return &tx, c.Call(ctx, &tx, methods.TxByHash, hash)
}

// BlockTxCountByHash returns pointer to transactions amount in block and call error if not nil.
func (c *Client) BlockTxCountByHash(ctx context.Context, hash common.Hash) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.BlockTxsCountByHash, hash)
}

// BlockTxCountByNumber returns pointer to transactions amount in block and call error if not nil.
func (c *Client) BlockTxCountByNumber(ctx context.Context, number *big.Int) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.BlockTxsCountByNumber, number)
}

// UncleCountByBlockHash returns pointer to uncles amount in block and call error if not nil.
func (c *Client) UncleCountByBlockHash(ctx context.Context, hash common.Hash) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.UncleCntByBlockHash, hash)
}

// UncleCountByBlockNumber returns pointer to uncles amount in block and call error if not nil.
func (c *Client) UncleCountByBlockNumber(ctx context.Context, number *big.Int) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.UncleCntByBlockNumber, number)
}

// TxByBlockHashAndIndex returns pointer to allocated and initialized models.Transaction and call error if not nil.
func (c *Client) TxByBlockHashAndIndex(ctx context.Context, hash common.Hash, index *big.Int) (*models.Transaction, error) {
	var res models.Transaction
	return &res, c.Call(ctx, &res, methods.TxByBlockHashAndIdx, hash, index)
}

// TxByBlockNumberAndIndex returns pointer to allocated and initialized models.Transaction and call error if not nil.
func (c *Client) TxByBlockNumberAndIndex(ctx context.Context, number *big.Int, index *big.Int) (*models.Transaction, error) {
	var res models.Transaction
	return &res, c.Call(ctx, &res, methods.TxByBlockNumberAndIdx, number, index)
}

// TxReceipt retrieves the transaction receipt identified by the given hash from the blockchain.
func (c *Client) TxReceipt(ctx context.Context, hash common.Hash) (*models.Receipt, error) {
	var res models.Receipt
	return &res, c.Call(ctx, &res, methods.TxReceipt, hash)
}

// UncleByBlockHashAndIndex retrieves the uncle block by its parent block hash and positional index in the block.
func (c *Client) UncleByBlockHashAndIndex(ctx context.Context, hash common.Hash, index *big.Int) (*models.Block, error) {
	var b models.Block
	return &b, c.Call(ctx, &b, methods.UncleByBlockHashAndIdx, hash, index)
}

// UncleByBlockNumberAndIndex retrieves an uncle block by its block number and index within the specified block.
func (c *Client) UncleByBlockNumberAndIndex(ctx context.Context, number *big.Int, index *big.Int) (res *models.Block, err error) {
	var b models.Block
	return &b, c.Call(ctx, &b, methods.UncleByBlockNumAndIdx, number, index)
}

// TxsCount retrieves the number of transactions sent from the given address at the specified block.
func (c *Client) TxsCount(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.TxsCount, address, block)
}

// Code retrieves the contract code at a specific address and block number from the blockchain.
func (c *Client) Code(ctx context.Context, address common.Address, block *big.Int) (*[]byte, error) {
	var cd models.Code
	return &cd.Value, c.Call(ctx, &cd, methods.Code, address, block)
}

// CallContract executes a smart contract call with the given address, data, and block, returning the result or an error.
func (c *Client) CallContract(ctx context.Context, address common.Address, data []byte, block *big.Int) (*[]byte, error) {
	var cd models.Code
	return &cd.Value, c.Call(ctx, &cd, methods.Call, map[string]interface{}{
		"to":   address.Hex(),
		"data": common.Bytes2Hex(data),
	}, block)
}

// EstimateGas estimates the gas needed to execute a given transaction without submitting it to the blockchain.
func (c *Client) EstimateGas(ctx context.Context, data []byte, block *big.Int) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.EstimateGas, data, block)
}

// Logs fetches logs for a given address, optional topics, and a specific block number within the blockchain system.
func (c *Client) Logs(ctx context.Context, address common.Address, topics [][]common.Hash, block *big.Int) (*models.Logs, error) {
	var res models.Logs
	return &res, c.Call(ctx, &res, methods.Logs, map[string]interface{}{
		"address": address.Hex(),
		"topics":  topics,
	}, block)
}

// NewFilter creates and installs a new filter with the specified arguments, returning the filter ID and an error if any.
func (c *Client) NewFilter(ctx context.Context, args map[string]interface{}) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.NewFilter, args)
}

// NewBlockFilter creates a new filter in the node for new block headers and returns the filter ID and an error if any.
func (c *Client) NewBlockFilter(ctx context.Context) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.NewBlockFilter)
}

// NewPendingTransactionFilter creates a new filter to monitor pending transactions in the Ethereum network.
func (c *Client) NewPendingTransactionFilter(ctx context.Context) (*big.Int, error) {
	var res Int
	return res.n, c.Call(ctx, &res, methods.NewPendingTransactionFilter)
}

// UninstallFilter removes a filter identified by the given ID from the client.
func (c *Client) UninstallFilter(ctx context.Context, id *big.Int) error {
	return c.Call(ctx, nil, methods.UninstallFilter, id)
}

// FilterChanges retrieves blockchain log changes for a given filter
func (c *Client) FilterChanges(ctx context.Context, id *big.Int) (*models.Logs, error) {
	var res models.Logs
	return &res, c.Call(ctx, &res, methods.FilterChanges, id)
}

// FilterLogs retrieves logs filtered by the specified ID using the provided context.
// It returns the filtered logs or an error if the operation fails.
func (c *Client) FilterLogs(ctx context.Context, id *big.Int) (*models.Logs, error) {
	var res models.Logs
	return &res, c.Call(ctx, &res, methods.FilterLogs, id)
}

// Sign signs the provided data using the private key associated with the specified address.
func (c *Client) Sign(ctx context.Context, address common.Address, data []byte) ([]byte, error) {
	var cd models.Code
	return cd.Value, c.Call(ctx, &cd, methods.Sign, map[string]interface{}{
		"address": address.Hex(),
		"data":    common.Bytes2Hex(data),
	})
}
