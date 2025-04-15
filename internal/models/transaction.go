package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/proto/extra"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type (
	extraTx struct {
		Data []byte
	}

	innerTx struct {
		BlockNumber *big.Int       `json:"blockNumber"`
		Value       *big.Int       `json:"value"`
		V           *big.Int       `json:"v"`
		R           *big.Int       `json:"r"`
		S           *big.Int       `json:"s"`
		Input       common.Hash    `json:"input"`
		Hash        common.Hash    `json:"hash"`
		From        common.Address `json:"from"`
		To          common.Address `json:"to"`
	}

	Transaction struct {
		inner innerTx
		extra extraTx
	}

	Transactions []Transaction
)

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

func (t *Transaction) GasPrice() (string, error) {
	var res extra.ExtraTx
	if err := proto.Unmarshal(t.extra.Data, &res); err != nil {
		return "", err
	}
	return res.GasPrice, nil
}

func (t *Transaction) Gas() (string, error) {
	var res extra.ExtraTx
	if err := proto.Unmarshal(t.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Gas, nil
}

func (t *Transaction) Nonce() (string, error) {
	var res extra.ExtraTx
	if err := proto.Unmarshal(t.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Nonce, nil
}

func (t *Transaction) TransactionIndex() (string, error) {
	var res extra.ExtraTx
	if err := proto.Unmarshal(t.extra.Data, &res); err != nil {
		return "", err
	}
	return res.TransactionIndex, nil
}

func (t *Transaction) BlockHash() (string, error) {
	var res extra.ExtraTx
	if err := proto.Unmarshal(t.extra.Data, &res); err != nil {
		return "", err
	}
	return res.BlockHash, nil
}
