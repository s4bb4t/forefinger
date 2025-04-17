package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"math/big"
)

type (
	Topics []common.Hash

	innerLog struct {
		Removed          bool
		Data             []byte
		Topics           Topics
		LogIndex         *big.Int
		TransactionIndex *big.Int
		BlockNumber      *big.Int
		TransactionHash  common.Hash
		Address          common.Hash
	}

	Log struct {
		inner innerLog
	}

	Logs []Log
)

func (l *Logs) Indirect() Logs {
	return *l
}

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
		case removed:
			l.inner.Removed = w.Bool()
		case data:
			l.inner.Data = w.Raw()
		case txIdx:
			l.inner.TransactionIndex.SetString(w.String(), 0)
		case logIdx:
			l.inner.LogIndex.SetString(w.String(), 0)
		case blockNum:
			l.inner.BlockNumber.SetString(w.String(), 0)
		case txHash:
			l.inner.TransactionHash = common.HexToHash(w.String())
		case address:
			l.inner.Address = common.HexToHash(w.String())
		case topics:
			l.inner.Topics.UnmarshalEasyJSON(w)
		default:
			w.SkipRecursive()
		}
		w.WantComma()
	}
	w.Delim('}')
}

func (t *Topics) UnmarshalEasyJSON(w *jlexer.Lexer) {
	w.Delim('[')
	for !w.IsDelim(']') {
		var hash common.Hash
		hash = common.HexToHash(w.String())
		*t = append(*t, hash)
		w.WantComma()
	}
	w.Delim(']')
}

func (l *Log) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, l)
}

func (l *Log) Removed() bool {
	return l.inner.Removed
}

func (l *Log) Data() []byte {
	return l.inner.Data
}

func (l *Log) TransactionIndex() *big.Int {
	return big.NewInt(0).Set(l.inner.TransactionIndex)
}

func (l *Log) LogIndex() *big.Int {
	return big.NewInt(0).Set(l.inner.LogIndex)
}

func (l *Log) BlockNumber() *big.Int {
	return big.NewInt(0).Set(l.inner.BlockNumber)
}

func (l *Log) TransactionHash() common.Hash {
	return l.inner.TransactionHash
}

func (l *Log) Address() common.Hash {
	return l.inner.Address
}

func (l *Log) Topics() Topics {
	return l.inner.Topics
}

func (l *Log) String() string {
	return string(l.inner.Data)
}

func (l *Log) Bytes() []byte {
	return l.inner.Data
}

func (l *Log) Hash() common.Hash {
	return common.BytesToHash(l.inner.Data)
}
