package methods

type Method string

const (
	BlockTxsCountByHash         Method = "eth_getBlockTransactionCountByHash"
	BlockTxsCountByNumber       Method = "eth_getBlockTransactionCountByNumber"
	UncleCntByBlockHash         Method = "eth_getUncleCountByBlockHash"
	UncleCntByBlockNumber       Method = "eth_getUncleCountByBlockNumber"
	BlockByHash                 Method = "eth_getBlockByHash"
	BlockByNumber               Method = "eth_getBlockByNumber"
	TxByHash                    Method = "eth_getTransactionByHash"
	TxByBlockHashAndIdx         Method = "eth_getTransactionByBlockHashAndIndex"
	TxByBlockNumberAndIdx       Method = "eth_getTransactionByBlockNumberAndIndex"
	TxReceipt                   Method = "eth_getTransactionReceipt"
	UncleByBlockHashAndIdx      Method = "eth_getUncleByBlockHashAndIndex"
	UncleByBlockNumAndIdx       Method = "eth_getUncleByBlockNumberAndIndex"
	Balance                     Method = "eth_getBalance"
	StorageAt                   Method = "eth_getStorageAt"
	TxsCount                    Method = "eth_getTransactionCount"
	Code                        Method = "eth_getCode"
	Call                        Method = "eth_call"
	EstimateGas                 Method = "eth_estimateGas"
	BlockNumber                        = "eth_blockNumber"
	Logs                        Method = "eth_getLogs"
	NewFilter                   Method = "eth_newFilter"
	NewBlockFilter              Method = "eth_newBlockFilter"
	NewPendingTransactionFilter Method = "eth_newPendingTransactionFilter"
	UninstallFilter             Method = "eth_uninstallFilter"
	FilterChanges               Method = "eth_getFilterChanges"
	FilterLogs                  Method = "eth_getFilterLogs"
	Sign                        Method = "eth_sign"
	SendRawTransaction          Method = "eth_sendRawTransaction"
	GetBadBlocks                Method = "debug_getBadBlocks"
	GetRawBlock                 Method = "debug_getRawBlock"
	Version                     Method = "net_version"
	Listening                   Method = "net_listening"
	PeerCount                   Method = "net_peerCount"
	GasPrice                    Method = "eth_gasPrice"
	Subscribe                   Method = "eth_subscribe"
	Unsubscribe                 Method = "eth_unsubscribe"
	Latest                             = "latest"
	Pending                            = "pending"
	Earliest                           = "earliest"
	Safe                               = "safe"
	Finalized                          = "finalized"
)

func (m Method) Method() string {
	return string(m)
}
