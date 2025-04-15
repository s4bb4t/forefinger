package models

import (
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/proto/extra"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type (
	extraBlock struct {
		Data []byte
	}

	inner struct {
		Transactions Transactions `json:"transactions"`
		Timestamp    *big.Int     `json:"timestamp"`
		Number       *big.Int     `json:"number"`
		Size         *big.Int     `json:"size"`
	}

	Block struct {
		extra extraBlock
		inner inner
	}
)

func (b *Block) UnmarshalEasyJSON(w *jlexer.Lexer) {
	var ex extra.ExtraBlock
	b.inner.Timestamp = big.NewInt(0)
	b.inner.Number = big.NewInt(0)
	b.inner.Size = big.NewInt(0)
	w.Delim('{')
	for !w.IsDelim('}') {
		key := w.String()
		w.WantColon()
		switch key {
		case txs:
			b.inner.Transactions.UnmarshalEasyJSON(w)
		case timestamp:
			b.inner.Timestamp.SetString(w.String(), 0)
		case size:
			b.inner.Size.SetString(w.String(), 0)
		case number:
			b.inner.Number.SetString(w.String(), 0)

		case gasUsed:
			ex.GasUsed = w.String()
		case gasLimit:
			ex.GasLimit = w.String()
		case diff:
			ex.Difficulty = w.String()
		case extraData:
			ex.ExtraData = w.String()
		case hash:
			ex.Hash = w.String()
		case nonce:
			ex.Nonce = w.String()
		case miner:
			ex.Miner = w.String()
		case stateRoot:
			ex.StateRoot = w.String()
		case receiptsRoot:
			ex.ReceiptsRoot = w.String()
		case txsRoot:
			ex.TransactionsRoot = w.String()
		case sha3Uncles:
			ex.Sha3Uncles = w.String()
		case parentHash:
			ex.ParentHash = w.String()
		case logsBloom:
			ex.LogsBloom = w.String()
		default:
			w.SkipRecursive()
		}
		w.WantComma()
	}
	d, err := proto.Marshal(&ex)
	if err != nil {
		w.AddError(fmt.Errorf("extraData marshaling error: %w", err))
	}
	b.extra.Data = d
	w.Delim('}')
}

func (b *Block) UnmarshalJSON(bytes []byte) error {
	if err := easyjson.Unmarshal(bytes, b); err != nil {
		return fmt.Errorf("forefinger: block data unmarshaling error: %w", err)
	}
	return nil
}

func (b *Block) Number() *big.Int {
	return big.NewInt(0).Set(b.inner.Number)
}

func (b *Block) Size() *big.Int {
	return big.NewInt(0).Set(b.inner.Size)
}

func (b *Block) Timestamp() *big.Int {
	return big.NewInt(0).Set(b.inner.Timestamp)
}

func (b *Block) Transactions() Transactions {
	return b.inner.Transactions
}

func (b *Block) ExtraData() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.ExtraData, nil
}

func (b *Block) Hash() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Hash, nil
}

func (b *Block) Miner() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Miner, nil
}

func (b *Block) Nonce() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Nonce, nil
}

func (b *Block) StateRoot() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
}

func (b *Block) ReceiptsRoot() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
}

func (b *Block) TransactionsRoot() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.TransactionsRoot, nil
}

func (b *Block) Sha3Uncles() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Sha3Uncles, nil
}

func (b *Block) ParentHash() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.ParentHash, nil
}

func (b *Block) LogsBloom() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.LogsBloom, nil
}

func (b *Block) Difficulty() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.Difficulty, nil
}

func (b *Block) GasLimit() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.GasLimit, nil
}

func (b *Block) GasUsed() (string, error) {
	var res extra.ExtraBlock
	if err := proto.Unmarshal(b.extra.Data, &res); err != nil {
		return "", err
	}
	return res.GasUsed, nil
}
