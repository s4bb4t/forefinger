package models

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/proto/extra"
	"google.golang.org/protobuf/proto"
	"math/big"
)

var exReceiptShared extra.ExtraReceipt

type (
	extraReceipt struct {
		Data []byte
	}
	innerReceipt struct {
		TransactionHash  common.Hash
		From             common.Address
		To               common.Address
		ContractAddress  common.Address
		TransactionIndex *big.Int
		BlockNumber      *big.Int
		Type             *big.Int
		Status           *big.Int
		Logs             Logs
	}
	Receipt struct {
		extra extraReceipt
		inner innerReceipt
	}

	Receipts []Receipt
)

func (r *Receipt) UnmarshalEasyJSON(w *jlexer.Lexer) {
	var ex extra.ExtraReceipt
	r.inner.TransactionIndex = big.NewInt(0)
	r.inner.BlockNumber = big.NewInt(0)
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
				ex.CumulativeGasUsed = w.String()
			}
		case effectiveGasPrice:
			ex.EffectiveGasPrice = w.String()
		case gasUsed:
			ex.GasUsed = w.String()
		case logsBloom:
			ex.LogsBloom = w.String()
		case root:
			ex.Root = w.String()

		case type_:
			r.inner.Type.SetString(w.String(), 0)
		case status:
			r.inner.Status.SetString(w.String(), 0)
		case from:
			r.inner.From = common.BytesToAddress(w.Raw())
		case to:
			r.inner.To = common.BytesToAddress(w.Raw())
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
	d, err := proto.Marshal(&ex)
	if err != nil {
		w.AddError(fmt.Errorf("extraData marshaling error: %w", err))
	}
	r.extra.Data = d
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

func (r *Receipt) CumulativeGasUsed() (*big.Int, error) {
	if err := proto.Unmarshal(r.extra.Data, &exReceiptShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exReceiptShared.CumulativeGasUsed, 0)
	if !ok {
		return nil, errors.New("failed to parse cumulative gas used")
	}
	return g, nil
}

func (r *Receipt) EffectiveGasPrice() (*big.Int, error) {
	if err := proto.Unmarshal(r.extra.Data, &exReceiptShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exReceiptShared.EffectiveGasPrice, 0)
	if !ok {
		return nil, errors.New("failed to parse cumulative gas used")
	}
	return g, nil
}

func (r *Receipt) GasUsed() (*big.Int, error) {
	if err := proto.Unmarshal(r.extra.Data, &exReceiptShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exReceiptShared.GasUsed, 0)
	if !ok {
		return nil, errors.New("failed to parse cumulative gas used")
	}
	return g, nil
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

func (r *Receipt) LogsBloom() (common.Hash, error) {
	if err := proto.Unmarshal(r.extra.Data, &exReceiptShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exReceiptShared.LogsBloom), nil
}

func (r *Receipt) Root() (common.Hash, error) {
	if err := proto.Unmarshal(r.extra.Data, &exReceiptShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exReceiptShared.Root), nil
}

func (r *Receipt) TransactionHash() common.Hash {
	return r.inner.TransactionHash
}

func (r *Receipt) Logs() Logs {
	return r.inner.Logs
}
