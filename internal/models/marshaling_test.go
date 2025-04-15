package models

import (
	"fmt"
	"math/big"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	n := big.NewInt(100)

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("0x%x", n.String())
	}
}

func BenchmarkConcatenate(b *testing.B) {
	n := big.NewInt(100)

	for i := 0; i < b.N; i++ {
		_ = "0x" + n.String()
	}
}
