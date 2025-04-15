package client

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/s4bb4t/forefinger/internal/models"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"math/big"
)

//BlockTxsCountByHash         Method = "eth_getBlockTransactionCountByHash"
//BlockTxsCountByNumber       Method = "eth_getBlockTransactionCountByNumber"
//UncleCntByBlockHash         Method = "eth_getUncleCountByBlockHash"
//UncleCntByBlockNumber       Method = "eth_getUncleCountByBlockNumber"
//BlockByHash                 Method = "eth_getBlockByHash"
//BlockByNumber               Method = "eth_getBlockByNumber"
//TxByHash                    Method = "eth_getTransactionByHash"
//TxByBlockHashAndIdx         Method = "eth_getTransactionByBlockHashAndIndex"
//TxByBlockNumberAndIdx       Method = "eth_getTransactionByBlockNumberAndIndex"
//TxReceipt                   Method = "eth_getTransactionReceipt"
//UncleByBlockHashAndIdx      Method = "eth_getUncleByBlockHashAndIndex"
//UncleByBlockNumAndIdx       Method = "eth_getUncleByBlockNumberAndIndex"
//Balance                     Method = "eth_getBalance"
//StorageAt                   Method = "eth_getStorageAt"
//TxsCount                    Method = "eth_getTransactionCount"
//Code                        Method = "eth_getCode"
//Call                        Method = "eth_call"
//EstimateGas                 Method = "eth_estimateGas"
//Logs                        Method = "eth_getLogs"
//NewFilter                   Method = "eth_newFilter"
//NewBlockFilter              Method = "eth_newBlockFilter"
//NewPendingTransactionFilter Method = "eth_newPendingTransactionFilter"
//UninstallFilter             Method = "eth_uninstallFilter"
//FilterChanges               Method = "eth_getFilterChanges"
//FilterLogs                  Method = "eth_getFilterLogs"
//Sign                        Method = "eth_sign"
//SendRawTransaction          Method = "eth_sendRawTransaction"
//GetBadBlocks                Method = "debug_getBadBlocks"
//GetRawBlock                 Method = "debug_getRawBlock"
//Version                     Method = "net_version"
//Listening                   Method = "net_listening"
//PeerCount                   Method = "net_peerCount"
//GasPrice                    Method = "eth_gasPrice"
//Subscribe                   Method = "eth_subscribe"
//Unsubscribe                 Method = "eth_unsubscribe"

func (c *Client) BlockByNumber(ctx context.Context, number *big.Int) (res *models.Block, err error) {
	return res, c.Call(ctx, res, methods.BlockByNumber, number, true)
}

func (c *Client) BlockByHash(ctx context.Context, hash common.Hash) (res *models.Block, err error) {
	return res, c.Call(ctx, res, methods.BlockByHash, hash, true)
}

func (c *Client) TransactionByHash(ctx context.Context, hash common.Hash) (res *models.Transaction, err error) {
	return res, c.Call(ctx, res, methods.TxByHash, hash)
}

func (c *Client) BlockTransactionCountByHash(ctx context.Context, hash common.Hash) (res *big.Int, err error) {
	return res, c.Call(ctx, &res, methods.BlockTxsCountByHash, hash)
}
func (c *Client) BlockTransactionCountByNumber(ctx context.Context, number *big.Int) (res *big.Int, err error) {
	return res, c.Call(ctx, &res, methods.BlockTxsCountByNumber, number)
}
func (c *Client) UncleCountByBlockHash(ctx context.Context, hash common.Hash) (res *big.Int, err error) {
	return res, c.Call(ctx, &res, methods.UncleCntByBlockHash, hash)
}
func (c *Client) UncleCountByBlockNumber(ctx context.Context, number *big.Int) (res *big.Int, err error) {
	return res, c.Call(ctx, &res, methods.UncleCntByBlockNumber, number)
}
func (c *Client) TxByBlockHashAndIndex(ctx context.Context, hash common.Hash, index *big.Int) (res *models.Transaction, err error) {
	return res, c.Call(ctx, res, methods.TxByBlockHashAndIdx, hash, index)
}
func (c *Client) TxByBlockNumberAndIndex(ctx context.Context, number *big.Int, index *big.Int) (res *models.Transaction, err error) {
	return res, c.Call(ctx, res, methods.TxByBlockNumberAndIdx, number, index)
}
func (c *Client) TxReceipt(ctx context.Context, hash common.Hash) (res *models.Receipt, err error) {
	return res, c.Call(ctx, res, methods.TxReceipt, hash)
}
func (c *Client) UncleByBlockHashAndIndex(ctx context.Context, hash common.Hash, index *big.Int) (res *models.Block, err error) {
	return res, c.Call(ctx, res, methods.UncleByBlockHashAndIdx, hash, index)
}
func (c *Client) UncleByBlockNumberAndIndex(ctx context.Context, number *big.Int, index *big.Int) (res *models.Block, err error) {
	return res, c.Call(ctx, res, methods.UncleByBlockNumAndIdx, number, index)
}
