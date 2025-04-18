package models

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
)

const (
	blockNum  = "blockNumber"
	gasPrice  = "gasPrice"
	gas       = "gas"
	nonce     = "nonce"
	txIdx     = "transactionIndex"
	val       = "value"
	v         = "v"
	r         = "r"
	s         = "s"
	blockHash = "blockHash"
	hash      = "hash"
	from      = "from"
	to        = "to"
	input     = "input"

	accessList           = "accessList"
	maxFeePerGas         = "maxFeePerGas"
	maxPriorityFeePerGas = "maxPriorityFeePerGas"
	maxFeePerBlobGas     = "maxFeePerBlobGas"
	blobVersionedHashes  = "blobVersionedHashes"
	beaconRoot           = "beaconRoot"
	chainId              = "chainId"

	txs          = "transactions"
	timestamp    = "timestamp"
	size         = "size"
	number       = "number"
	gasUsed      = "gasUsed"
	gasLimit     = "gasLimit"
	diff         = "difficulty"
	extraData    = "extraData"
	data         = "data"
	miner        = "miner"
	stateRoot    = "stateRoot"
	receiptsRoot = "receiptsRoot"
	txsRoot      = "transactionsRoot"
	sha3Uncles   = "sha3Uncles"
	parentHash   = "parentHash"
	logsBloom    = "logsBloom"

	cumulativeGasUsed = "cumulativeGasUsed"
	effectiveGasPrice = "effectiveGasPrice"
	type_             = "type"
	status            = "status"
	root              = "root"
	contractAddress   = "contractAddress"
	logs              = "logs"

	removed = "removed"
	topics  = "topics"
	logIdx  = "logIndex"
	txHash  = "transactionHash"
	address = "address"
)

type Code struct {
	Value []byte
}

func (c *Code) UnmarshalEasyJSON(w *jlexer.Lexer) {
	c.Value = w.Raw()
}

func (c *Code) UnmarshalJSON(b []byte) error {
	return easyjson.Unmarshal(b, c)
}
