package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/proto/extra"
	"google.golang.org/protobuf/proto"
	"math/big"
)

var exTxShared extra.ExtraTx

type (
	extraTx struct {
		Data []byte
	}

	innerTx struct {
		BlockNumber *big.Int
		Value       *big.Int
		V           *big.Int
		R           *big.Int
		S           *big.Int
		Input       common.Hash
		Hash        common.Hash
		From        common.Address
		To          common.Address
	}

	Transaction struct {
		inner innerTx
		extra extraTx
	}

	Transactions []Transaction
)

func (t *Transactions) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, t)
}

func (t *Transaction) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, t)
}

func (t *Transactions) UnmarshalEasyJSON(w *jlexer.Lexer) {
	w.Delim('[')
	for !w.IsDelim(']') {
		var tx Transaction
		tx.UnmarshalEasyJSON(w)
		*t = append(*t, tx)
		w.WantComma()
	}
	w.Delim(']')
}

func (t *Transaction) UnmarshalEasyJSON(w *jlexer.Lexer) {
	var ex extra.ExtraTx
	t.inner.BlockNumber = big.NewInt(0)
	t.inner.Value = big.NewInt(0)
	t.inner.V = big.NewInt(0)
	t.inner.R = big.NewInt(0)
	t.inner.S = big.NewInt(0)
	w.Delim('{')
	for !w.IsDelim('}') {
		key := w.String()
		w.WantColon()
		switch key {
		case gasPrice:
			ex.GasPrice = w.String()
		case txIdx:
			ex.TransactionIndex = w.String()
		case nonce:
			ex.Nonce = w.String()
		case gas:
			ex.Gas = w.String()
		case blockHash:
			ex.BlockHash = w.String()

		case blockNum:
			t.inner.BlockNumber.SetString(w.String(), 0)
		case val:
			t.inner.Value.SetString(w.String(), 0)
		case v:
			t.inner.V.SetString(w.String(), 0)
		case r:
			t.inner.R.SetString(w.String(), 0)
		case s:
			t.inner.S.SetString(w.String(), 0)

		case hash:
			t.inner.Hash = common.HexToHash(w.String())
		case from:
			t.inner.From = common.HexToAddress(w.String())
		case to:
			t.inner.To = common.HexToAddress(w.String())
		case input:
			t.inner.Input = common.HexToHash(w.String())
		default:
			w.SkipRecursive()
		}
		w.WantComma()
	}
	d, err := proto.Marshal(&ex)
	if err != nil {
		w.AddError(fmt.Errorf("extraData marshaling error: %w", err))
	}
	t.extra.Data = d
	w.Delim('}')
}

// BlockNumber returns the block number of the transaction as a *big.Int.
func (t *Transaction) BlockNumber() *big.Int {
	return big.NewInt(0).Set(t.inner.BlockNumber)
}

// Value returns the transaction value as a pointer to a big.Int.
func (t *Transaction) Value() *big.Int {
	return big.NewInt(0).Set(t.inner.Value)
}

// V returns the 'V' value of the transaction as a pointer to a big.Int.
func (t *Transaction) V() *big.Int {
	return big.NewInt(0).Set(t.inner.V)
}

// R returns the 'R' value of the transaction as a pointer to a big.Int.
func (t *Transaction) R() *big.Int {
	return big.NewInt(0).Set(t.inner.R)
}

// S returns the 'S' value of the transaction as a pointer to a big.Int.
func (t *Transaction) S() *big.Int {
	return big.NewInt(0).Set(t.inner.S)
}

// Input returns the 'Input' field of the transaction as a common.Hash.
func (t *Transaction) Input() common.Hash {
	return t.inner.Input
}

// Hash returns the hash of the transaction as a common.Hash.
func (t *Transaction) Hash() common.Hash {
	return t.inner.Hash
}

// From returns the sender's address of the transaction as a common.Hash.
func (t *Transaction) From() common.Address {
	return t.inner.From
}

// To returns the recipient address of the transaction as a common.Hash.
func (t *Transaction) To() common.Address {
	return t.inner.To
}

func (t *Transaction) GasPrice() (*big.Int, error) {
	if err := proto.Unmarshal(t.extra.Data, &exTxShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exTxShared.GasPrice, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse gas price")
	}
	return g, nil
}

func (t *Transaction) Gas() (*big.Int, error) {
	if err := proto.Unmarshal(t.extra.Data, &exTxShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exTxShared.Gas, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse gas")
	}
	return g, nil
}

func (t *Transaction) Nonce() (*big.Int, error) {
	if err := proto.Unmarshal(t.extra.Data, &exTxShared); err != nil {
		return nil, err
	}
	n, ok := big.NewInt(0).SetString(exTxShared.Nonce, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse nonce")
	}
	return n, nil
}

func (t *Transaction) TransactionIndex() (*big.Int, error) {
	if err := proto.Unmarshal(t.extra.Data, &exTxShared); err != nil {
		return nil, err
	}
	i, ok := big.NewInt(0).SetString(exTxShared.TransactionIndex, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse transaction index")
	}
	return i, nil
}

func (t *Transaction) BlockHash() (common.Hash, error) {
	if err := proto.Unmarshal(t.extra.Data, &exTxShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exTxShared.BlockHash), nil
}
