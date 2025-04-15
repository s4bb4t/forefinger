package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson/jlexer"
	"math/big"
)

type (
	innerLog struct {
		Removed          bool          `json:"removed"`
		Data             []byte        `json:"data"`
		Topics           []common.Hash `json:"topics"`
		LogIndex         *big.Int      `json:"logIndex"`
		TransactionIndex *big.Int      `json:"transactionIndex"`
		BlockNumber      *big.Int      `json:"blockNumber"`
		TransactionHash  common.Hash   `json:"transactionHash"`
		Address          common.Hash   `json:"address"`
	}

	Log struct {
		inner innerLog
	}

	Logs []Log
)

func (l *Logs) UnmarshalEasyJSON(w *jlexer.Lexer) {
	w.Delim('[')
	for !w.IsDelim(']') {
		var log Log
		log.UnmarshalEasyJSON(w)
		*l = append(*l, log)
		w.WantComma()
	}
	w.Delim(']')
}

func (l *Log) UnmarshalEasyJSON(w *jlexer.Lexer) {
	l.inner.BlockNumber = big.NewInt(0)
	l.inner.LogIndex = big.NewInt(0)
	l.inner.TransactionIndex = big.NewInt(0)
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
			t.inner.From = common.HexToHash(w.String())
		case to:
			t.inner.To = common.HexToHash(w.String())
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

func (l *Log) UnmarshalJSON(bytes []byte) error {
	return nil
}
