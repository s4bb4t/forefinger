package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"math/big"
)

type (
	innerReceipt struct {
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
		Logs              Logs
	}
	Receipt struct {
		inner innerReceipt
	}

	Receipts []Receipt
)

func (r *Receipt) UnmarshalEasyJSON(w *jlexer.Lexer) {
	r.inner.TransactionIndex = big.NewInt(0)
	r.inner.BlockNumber = big.NewInt(0)
	r.inner.CumulativeGasUsed = big.NewInt(0)
	r.inner.EffectiveGasPrice = big.NewInt(0)
	r.inner.GasUsed = big.NewInt(0)
	r.inner.Type = big.NewInt(0)
	r.inner.Status = big.NewInt(0)
	w.Delim('{')
	for !w.IsDelim('}') {
		key := w.String()
		w.WantColon()
		switch key {
		case txIdx:
			r.inner.TransactionIndex.SetString(w.String(), 0)
		case blockNum:
			r.inner.BlockNumber.SetString(w.String(), 0)
		case cumulativeGasUsed:
			if w.IsNull() {
				w.Skip()
			} else {
				r.inner.CumulativeGasUsed.SetString(w.String(), 0)
			}
		case effectiveGasPrice:
			r.inner.EffectiveGasPrice.SetString(w.String(), 0)
		case gasUsed:
			r.inner.GasUsed.SetString(w.String(), 0)
		case type_:
			r.inner.Type.SetString(w.String(), 0)
		case status:
			r.inner.Status.SetString(w.String(), 0)

		case from:
			r.inner.From = common.BytesToAddress(w.Raw())
		case to:
			r.inner.To = common.BytesToAddress(w.Raw())
		case logsBloom:
			r.inner.LogsBloom = common.BytesToHash(w.Raw())
		case root:
			r.inner.Root = common.BytesToHash(w.Raw())
		case txHash:
			r.inner.TransactionHash = common.BytesToHash(w.Raw())
		case logs:
			r.inner.Logs.UnmarshalEasyJSON(w)
		case contractAddress:
			if w.IsNull() {
				w.Skip()
				r.inner.ContractAddress = common.Address{}
			} else {
				r.inner.ContractAddress = common.HexToAddress(w.String())
			}
		default:
			w.SkipRecursive()
		}
		w.WantComma()
	}
	w.Delim('}')
}

func (r *Receipts) UnmarshalEasyJSON(w *jlexer.Lexer) {
	w.Delim('[')
	for !w.IsDelim(']') {
		var res Receipt
		res.UnmarshalEasyJSON(w)
		*r = append(*r, res)
		w.WantComma()
	}
	w.Delim(']')
}

func (r *Receipt) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, r)
}

func (r *Receipts) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, r)
}

func (r *Receipt) TransactionIndex() *big.Int {
	return big.NewInt(0).Set(r.inner.TransactionIndex)
}

func (r *Receipt) BlockNumber() *big.Int {
	return big.NewInt(0).Set(r.inner.BlockNumber)
}

func (r *Receipt) CumulativeGasUsed() *big.Int {
	return big.NewInt(0).Set(r.inner.CumulativeGasUsed)
}

func (r *Receipt) EffectiveGasPrice() *big.Int {
	return big.NewInt(0).Set(r.inner.EffectiveGasPrice)
}

func (r *Receipt) GasUsed() *big.Int {
	return big.NewInt(0).Set(r.inner.GasUsed)
}

func (r *Receipt) Type() *big.Int {
	return big.NewInt(0).Set(r.inner.Type)
}

func (r *Receipt) Status() *big.Int {
	return big.NewInt(0).Set(r.inner.Status)
}

func (r *Receipt) From() common.Address {
	return r.inner.From
}

func (r *Receipt) To() common.Address {
	return r.inner.To
}

func (r *Receipt) ContractAddress() common.Address {
	return r.inner.ContractAddress
}

func (r *Receipt) LogsBloom() common.Hash {
	return r.inner.LogsBloom
}

func (r *Receipt) Root() common.Hash {
	return r.inner.Root
}

func (r *Receipt) TransactionHash() common.Hash {
	return r.inner.TransactionHash
}

func (r *Receipt) Logs() Logs {
	return r.inner.Logs
}
