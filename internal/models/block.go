package models

type (
	info struct {
		// Pointer embedded fields
		uncles           *[][32]byte
		logsBloom        *[256]byte
		parentHash       *[32]byte
		difficultyHash   *[32]byte
		sha3Uncles       *[32]byte
		transactionsRoot *[32]byte
		receiptsRoot     *[32]byte
		stateRoot        *[32]byte
		miner            *[20]byte
		nonce            *[8]byte
		gasUsed          *Quantity
		gasLimit         *Quantity
		difficulty       *Quantity
		extraData        *[]byte
	}

	Block struct {
		// Value embedded fields
		transactions []Transaction
		timestamp    Quantity
		size         Quantity
		number       Quantity
		hash         [32]byte

		info
	}
)

func NewBlock() *Block {
	return &Block{}
}

func (b *Block) Copy(to *Block) *Block {
	*to = *b
	return to
}

// Randomize sets random data to `b` included embedded fields
func (b *Block) Randomize() *Block {
	return b
}

// Transactions returns copy of Block's transactions
func (b *Block) Transactions() []Transaction {
	return b.transactions
}

// Timestamp returns copy of Block's timestamp
func (b *Block) Timestamp() SInt {
	return SInt(b.timestamp.Load())
}

// Size returns copy of Block's size in bytes
func (b *Block) Size() SInt {
	return SInt(b.size.Load())
}

// Number returns copy of Block's number in blockchain
func (b *Block) Number() SInt {
	return SInt(b.number.Load())
}

// Hash returns copy of Block's hash
func (b *Block) Hash() [32]byte {
	return b.hash
}

func (b *Block) Uncles() *[][32]byte {
	return b.uncles
}

func (b *Block) LogsBloom() *[256]byte {
	return b.logsBloom
}

func (b *Block) ParentHash() *[32]byte {
	return b.parentHash
}

func (b *Block) Sha3Uncles() *[32]byte {
	return b.sha3Uncles
}

func (b *Block) TransactionsRoot() *[32]byte {
	return b.transactionsRoot
}

func (b *Block) ReceiptsRoot() *[32]byte {
	return b.receiptsRoot
}

func (b *Block) StateRoot() *[32]byte {
	return b.stateRoot
}

func (b *Block) Miner() *[20]byte {
	return b.miner
}

func (b *Block) Nonce() *[8]byte {
	return b.nonce
}

func (b *Block) GasUsed() *Quantity {
	return b.gasUsed
}

func (b *Block) Difficulty() *Quantity {
	return b.difficulty
}

func (b *Block) ExtraData() *[]byte {
	return b.extraData
}

func (b *Block) GasLimit() *Quantity {
	return b.gasLimit
}

func (b *Block) DifficultyHash() *[32]byte {
	return b.difficultyHash
}
