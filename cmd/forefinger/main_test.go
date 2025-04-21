package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/s4bb4t/forefinger/pkg/client"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"github.com/s4bb4t/forefinger/pkg/models"
	"math/big"
	"sync"
	"testing"
)

var l = 2000

func TestSingle(t *testing.T) {
	cl, err := client.NewClient("http://10.255.13.100:8545", 100)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < l; i++ {
		block, err := cl.BlockByNumber(context.Background(), big.NewInt(22000000+int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		_ = block
	}
}

func TestSingleGo(t *testing.T) {
	cl, err := client.NewClient("http://10.255.13.100:8545", 100)
	if err != nil {
		t.Fatal(err)
	}
	blocks := make(chan *models.Block, l)
	var wg sync.WaitGroup

	for i := 0; i < l; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			block, err := cl.BlockByNumber(context.Background(), big.NewInt(22000000+int64(i)))
			if err != nil {
				t.Log(err)
			}
			blocks <- block
		}()
	}

	wg.Wait()
	close(blocks)
}

func TestBatch(t *testing.T) {
	cl, err := client.NewClient("http://10.255.13.100:8545", 100)
	if err != nil {
		t.Fatal(err)
	}

	blocks := make([]any, l)
	args := make([][]any, l)
	for i := 0; i < l; i++ {
		blocks[i] = models.Block{}
		args[i] = []any{big.NewInt(22000000 + int64(i)), true}
	}

	err, _ = cl.BatchCall(context.Background(), 100, methods.BlockByNumber, blocks, args)
	if err != nil {
		t.Fatal(err)
	}
}

func Benchmark_GoEthereum_Single(b *testing.B) {
	cl, err := ethclient.Dial("http://10.255.13.100:8545")
	if err != nil {
		b.Fatalf("failed to create go-ethereum client: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		block, err := cl.BlockByNumber(context.Background(), big.NewInt(22000000))
		if err != nil {
			b.Fatalf("failed to fetch block using go-ethereum client: %v", err)
		}
		_ = block
	}
}

func BenchmarkForefinger_Single(b *testing.B) {
	cl, err := client.NewClient("http://10.255.13.100:8545", 100)
	if err != nil {
		fmt.Println(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var block models.Block
		if err := cl.Call(context.Background(), &block, methods.BlockByNumber, big.NewInt(22000000), true); err != nil {
			b.Fatalf("failed to fetch block using custom solution: %v", err)
		}
	}
}

func Benchmark_GoEthereum_Cycle(b *testing.B) {
	cl, err := ethclient.Dial("http://10.255.13.100:8545")
	if err != nil {
		b.Fatalf("failed to create go-ethereum client: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < 100; i++ {
		b.N++
		block, err := cl.BlockByNumber(context.Background(), big.NewInt(22000000))
		if err != nil {
			b.Fatalf("failed to fetch block using go-ethereum client: %v", err)
		}
		_ = block
	}
}

//func Benchmark_Forefinger_Batch(b *testing.B) {
//	cl, err := client.NewClient("http://10.255.13.100:8545", 100)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	r := make([]any, 100)
//	args := make([][]any, 100)
//
//	for i := 0; i < 100; i++ {
//		b.N++
//		args[i] = []any{big.NewInt(22000000), true}
//	}
//
//	b.ResetTimer()
//	if err, errs := cl.BatchCall(context.Background(), 50, methods.BlockByNumber, &r, args); err != nil {
//		fmt.Println(err, errs)
//	}
//}
