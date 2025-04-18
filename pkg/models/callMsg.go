package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type CallMsg struct {
	from      common.Address
	to        *common.Address
	gas       uint64
	gasPrice  *big.Int
	gasFeeCap *big.Int
	gasTipCap *big.Int
	value     *big.Int
	data      []byte

	accessList types.AccessList

	blobGasFeeCap *big.Int
	blobHashes    []common.Hash
}

func NewCallMsg() *CallMsg {
	return &CallMsg{}
}

func (m *CallMsg) From(addr common.Address) *CallMsg {
	m.from = addr
	return m
}

func (m *CallMsg) To(addr common.Address) *CallMsg {
	m.to = &addr
	return m
}

func (m *CallMsg) Gas(gas uint64) *CallMsg {
	m.gas = gas
	return m
}

func (m *CallMsg) GasPrice(gasPrice *big.Int) *CallMsg {
	m.gasPrice = gasPrice
	return m
}

func (m *CallMsg) GasFeeCap(gasFeeCap *big.Int) *CallMsg {
	m.gasFeeCap = gasFeeCap
	return m
}

func (m *CallMsg) GasTipCap(gasTipCap *big.Int) *CallMsg {
	m.gasTipCap = gasTipCap
	return m
}

func (m *CallMsg) Value(value *big.Int) *CallMsg {
	m.value = value
	return m
}

func (m *CallMsg) Data(data []byte) *CallMsg {
	m.data = data
	return m
}

func (m *CallMsg) AccessList(accessList types.AccessList) *CallMsg {
	m.accessList = accessList
	return m
}

func (m *CallMsg) BlobGasFeeCap(blobGasFeeCap *big.Int) *CallMsg {
	m.blobGasFeeCap = blobGasFeeCap
	return m
}

func (m *CallMsg) BlobHashes(blobHashes []common.Hash) *CallMsg {
	m.blobHashes = blobHashes
	return m
}

func (m *CallMsg) ToCallArg() interface{} {
	arg := map[string]interface{}{
		"from": m.from,
		"to":   m.to,
	}
	if len(m.data) > 0 {
		arg["input"] = hexutil.Bytes(m.data)
	}
	if m.value != nil {
		arg["value"] = (*hexutil.Big)(m.value)
	}
	if m.gas != 0 {
		arg["gas"] = hexutil.Uint64(m.gas)
	}
	if m.gasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(m.gasPrice)
	}
	if m.gasFeeCap != nil {
		arg["maxFeePerGas"] = (*hexutil.Big)(m.gasFeeCap)
	}
	if m.gasTipCap != nil {
		arg["maxPriorityFeePerGas"] = (*hexutil.Big)(m.gasTipCap)
	}
	if m.accessList != nil {
		arg["accessList"] = m.accessList
	}
	if m.blobGasFeeCap != nil {
		arg["maxFeePerBlobGas"] = (*hexutil.Big)(m.blobGasFeeCap)
	}
	if m.blobHashes != nil {
		arg["blobVersionedHashes"] = m.blobHashes
	}
	return arg
}
