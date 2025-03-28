package models

/*
	blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
	gas: QUANTITY - gas provided by the sender.
	gasPrice: QUANTITY - gas price provided by the sender in Wei.
	nonce: QUANTITY - the number of transactions made by the sender prior to this one.
	transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
	value: QUANTITY - value transferred in Wei.
	v: QUANTITY - ECDSA recovery id
	r: QUANTITY - ECDSA signature r
	s: QUANTITY - ECDSA signature s
*/

type Transaction struct {
	gasPrice         *Quantity
	gas              *Quantity
	nonce            *Quantity
	transactionIndex *Quantity
	value            *Quantity
	v                *Quantity
	r                *Quantity
	s                *Quantity
	input            []byte
	hash             [32]byte
	blockHash        [32]byte
	from             [20]byte
	to               [20]byte
}

func NewTransaction() *Transaction {
	return &Transaction{
		gasPrice:         NewQuantity(),
		gas:              NewQuantity(),
		nonce:            NewQuantity(),
		transactionIndex: NewQuantity(),
		value:            NewQuantity(),
		v:                NewQuantity(),
		r:                NewQuantity(),
		s:                NewQuantity(),
		input:            nil,
		hash:             [32]byte{},
		blockHash:        [32]byte{},
		from:             [20]byte{},
		to:               [20]byte{},
	}
}

func (t *Transaction) Randomize() *Transaction {
	return t
}
