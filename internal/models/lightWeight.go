package models

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	mathRand "math/rand"
)

type (
	toCompress struct {
		Uncles           [][32]byte `json:"uncles"`
		LogsBloom        [256]byte  `json:"logs_bloom"`
		ParentHash       [32]byte   `json:"parent_hash"`
		DifficultyHash   [32]byte   `json:"difficulty_hash"`
		Sha3Uncles       [32]byte   `json:"sha_3_uncles"`
		TransactionsRoot [32]byte   `json:"transactions_root"`
		ReceiptsRoot     [32]byte   `json:"receipts_root"`
		StateRoot        [32]byte   `json:"state_root"`
		Miner            [20]byte   `json:"miner"`
		Nonce            [8]byte    `json:"nonce"`
		GasUsed          []byte     `json:"gas_used"`
		GasLimit         []byte     `json:"gas_limit"`
		Difficulty       []byte     `json:"difficulty"`
		ExtraData        []byte     `json:"extra_data"`
	}

	LightWeightBlock struct {
		hash         [32]byte
		transactions []Transaction
		timestamp    Quantity
		size         Quantity
		number       Quantity

		compressed *[]byte
	}
)

func NewLightWeightBlock() *LightWeightBlock {
	return &LightWeightBlock{
		hash:         [32]byte{},
		transactions: make([]Transaction, 0),
		timestamp:    *NewQuantity(),
		size:         *NewQuantity(),
		number:       *NewQuantity(),
		compressed:   &[]byte{},
	}
}

func newToCompress() *toCompress {
	return &toCompress{}
}

func (c *toCompress) Randomize() *toCompress {
	c.Uncles = make([][32]byte, mathRand.Int())
	for i := range c.Uncles {
		c.Uncles[i] = randHash()
	}
	salt := make([]byte, 256)
	rand.Read(salt)

	c.LogsBloom = [256]byte(salt)
	c.ParentHash = randHash()
	c.Sha3Uncles = randHash()
	c.TransactionsRoot = randHash()
	c.ReceiptsRoot = randHash()
	c.StateRoot = randHash()
	c.DifficultyHash = randHash()

	return c
}

func randHash() [32]byte {
	salt := make([]byte, 32)
	rand.Read(salt)

	return [32]byte(salt)
}

// TODO: look for better variants of compressing
func (c *toCompress) compress(dest *[]byte) error {
	var err error
	*dest, err = json.Marshal(*c)
	if err != nil {
		return fmt.Errorf("error marshalling compressed block data: %w", err)
	}

	return nil
}

// Randomize sets random data to `b` included embedded fields
func (b *LightWeightBlock) Randomize() (*LightWeightBlock, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	b.hash = [32]byte(salt)

	b.timestamp = *NewQuantity().Set(mathRand.Uint64())
	b.size = *NewQuantity().Set(mathRand.Uint64())
	b.number = *NewQuantity().Set(mathRand.Uint64())

	b.transactions = make([]Transaction, 0)

	err := newToCompress().compress(b.compressed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate: %w", err)
	}

	return b, nil
}

// Transactions returns copy of LightWeightBlock's transactions
func (b *LightWeightBlock) Transactions() []Transaction {
	return b.transactions
}

// Timestamp returns copy of LightWeightBlock's timestamp
func (b *LightWeightBlock) Timestamp() SInt {
	return SInt(b.timestamp.Load())
}

// Size returns copy of LightWeightBlock's size in bytes
func (b *LightWeightBlock) Size() SInt {
	return SInt(b.size.Load())
}

// Number returns copy of LightWeightBlock's number in blockchain
func (b *LightWeightBlock) Number() SInt {
	return SInt(b.number.Load())
}

// Hash returns copy of LightWeightBlock's hash
func (b *LightWeightBlock) Hash() [32]byte {
	return b.hash
}
