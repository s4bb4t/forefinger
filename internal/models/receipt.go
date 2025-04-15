package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//transactionHash : DATA, 32 Bytes - hash of the transaction.
//transactionIndex: QUANTITY - integer of the transactions index position in the block.
//blockHash: DATA, 32 Bytes - hash of the block where this transaction was in.
//blockNumber: QUANTITY - block number where this transaction was in.
//from: DATA, 20 Bytes - address of the sender.
//to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//cumulativeGasUsed : QUANTITY - The total amount of gas used when this transaction was executed in the block.
//effectiveGasPrice : QUANTITY - The sum of the base fee and tip paid per unit of gas.
//gasUsed : QUANTITY - The amount of gas used by this specific transaction alone.
//contractAddress : DATA, 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise null.
//logs: Array - Array of log objects, which this transaction generated.
//logsBloom: DATA, 256 Bytes - Bloom filter for light clients to quickly retrieve related logs.
//type: QUANTITY - integer of the transaction type, 0x0 for legacy transactions, 0x1 for access list types, 0x2 for dynamic fees.
//root : DATA 32 bytes of post-transaction stateroot (pre Byzantium)
//status: QUANTITY either 1 (success) or 0 (failure)

type Receipt struct {
	TransactionHash   common.Hash
	LogsBloom         common.Hash
	Root              common.Hash
	From              common.Address
	To                common.Address
	ContractAddress   common.Address
	TransactionIndex  *big.Int
	BlockNumber       *big.Int
	CumulativeGasUsed *big.Int
	EffectiveGasPrice *big.Int
	GasUsed           *big.Int
	Type              *big.Int
	Status            *big.Int
	Logs              []Log
}
