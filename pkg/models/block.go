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

var exBlockShared extra.ExtraBlock

type (
	extraBlock struct {
		Data []byte
	}

	inner struct {
		Transactions Transactions
		Timestamp    *big.Int
		Number       *big.Int
		Size         *big.Int
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
	w.Delim('}')
	d, err := proto.Marshal(&ex)
	if err != nil {
		w.AddError(fmt.Errorf("extraData marshaling error: %w", err))
	}
	b.extra.Data = d
}

func (b *Block) UnmarshalJSON(bytes []byte) error {
	return easyjson.Unmarshal(bytes, b)
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

func (b *Block) ExtraData() ([]byte, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return nil, err
	}
	return []byte(exBlockShared.ExtraData), nil
}

func (b *Block) Hash() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.Hash), nil
}

func (b *Block) Miner() (common.Address, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Address{}, err
	}
	return common.HexToAddress(exBlockShared.Miner), nil
}

func (b *Block) Nonce() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.Nonce), nil
}

func (b *Block) StateRoot() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.StateRoot), nil
}

func (b *Block) ReceiptsRoot() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.ReceiptsRoot), nil
}

func (b *Block) TxsRoot() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.TransactionsRoot), nil
}

func (b *Block) Sha3Uncles() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.Sha3Uncles), nil
}

func (b *Block) ParentHash() (common.Hash, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(exBlockShared.ParentHash), nil
}

func (b *Block) Difficulty() (*big.Int, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exBlockShared.Difficulty, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse difficulty")
	}
	return g, nil
}

func (b *Block) GasLimit() (*big.Int, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exBlockShared.GasLimit, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse difficulty")
	}
	return g, nil
}

func (b *Block) GasUsed() (*big.Int, error) {
	if err := proto.Unmarshal(b.extra.Data, &exBlockShared); err != nil {
		return nil, err
	}
	g, ok := big.NewInt(0).SetString(exBlockShared.GasUsed, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse difficulty")
	}
	return g, nil
}
