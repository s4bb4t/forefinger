package models

import "sync/atomic"

const (
	latestS    string = "latest"
	pendingS   string = "pending"
	safeS      string = "safe"
	finalizedS string = "finalized"

	null      uint8 = 0
	latest    uint8 = 1
	pending   uint8 = 2
	safe      uint8 = 3
	finalized uint8 = 4
)

type Quantity struct {
	atomic.Uint64
}

func NewQuantity() *Quantity {
	return &Quantity{}
}

func (q *Quantity) Set(val uint64) *Quantity {
	q.Store(val)
	return q
}
