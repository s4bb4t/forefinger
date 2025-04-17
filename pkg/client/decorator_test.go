package client

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const (
	testAddressHex = "0x3fBBA2b0e07895ae8638D17fd83d72338954D272"
	testBlockHash  = "0xdf29aafca34c510304dffd0182ea648afe84efd463f9752bfb35de71d3423b37"
	testTxHashs    = "0xc48d5e8a0f2746cd885d9246fbfa083e5bd435c607ddbacd4680863ab66deb84"
	latest         = "latest"
	blockNumber    = 22281939
)

var (
	testAddress = common.HexToAddress(testAddressHex)
	testHash    = common.HexToHash(testBlockHash)
	testTxHash  = common.HexToHash(testTxHashs)
)

var cl, _ = NewClient("http://10.255.13.100:8545", 100)

func TestBlockByNumber(t *testing.T) {
	ctx := context.Background()
	blockNumber := big.NewInt(blockNumber)

	block, err := cl.BlockByNumber(ctx, blockNumber)
	if err != nil {
		t.Fatalf("BlockByNumber failed: %v", err)
	}

	if block == nil {
		t.Fatal("BlockByNumber returned nil block")
	}

	if block.Number().Int64() != blockNumber.Int64() {
		t.Fatalf("BlockByNumber returned wrong block number: %v", block.Number().Int64())
	}
	t.Log(block.Number().Int64())
}

func TestBalance(t *testing.T) {
	ctx := context.Background()

	balance, err := cl.Balance(ctx, testAddress, latest)
	if err != nil {
		t.Fatalf("Balance failed: %v", err)
	}

	if balance == nil {
		t.Fatal("Balance returned nil value")
	}

	t.Logf("Balance: %v", balance.String())
}

func TestBlockByHash(t *testing.T) {
	ctx := context.Background()

	block, err := cl.BlockByHash(ctx, testHash)
	if err != nil {
		t.Fatalf("BlockByHash failed: %v", err)
	}

	if block == nil {
		t.Fatal("BlockByHash returned nil block")
	}

	if block.Number().Int64() != blockNumber {
	}

	t.Logf("BlockByHash: %+v", block.Number().String())
}

func TestTxByHash(t *testing.T) {
	ctx := context.Background()

	tx, err := cl.TxByHash(ctx, testTxHash)
	if err != nil {
		t.Fatalf("TxByHash failed: %v", err)
	}

	if tx == nil {
		t.Fatal("TxByHash returned nil transaction")
	}

	if tx.Hash().String() != testTxHash.String() {
		t.Fatalf("TxByHash returned wrong transaction hash: %v", tx.Hash().String())
	}

	t.Logf("TxByHash: %+v", tx.Hash().String())
}

func TestBlockTxCountByHash(t *testing.T) {
	ctx := context.Background()

	count, err := cl.BlockTxCountByHash(ctx, testHash)
	if err != nil {
		t.Fatalf("BlockTxCountByHash failed: %v", err)
	}

	if count == nil {
		t.Fatal("BlockTxCountByHash returned nil value")
	}

	t.Logf("BlockTxCountByHash: %v", count.String())
}

func TestUncleCountByBlockHash(t *testing.T) {
	ctx := context.Background()

	count, err := cl.UncleCountByBlockHash(ctx, testHash)
	if err != nil {
		t.Fatalf("UncleCountByBlockHash failed: %v", err)
	}

	if count == nil {
		t.Fatal("UncleCountByBlockHash returned nil value")
	}

	t.Logf("UncleCountByBlockHash: %v", count.String())
}

func TestTxReceipt(t *testing.T) {
	ctx := context.Background()

	receipt, err := cl.TxReceipt(ctx, testTxHash)
	if err != nil {
		t.Fatalf("TxReceipt failed: %v", err)
	}

	if receipt == nil {
		t.Fatal("TxReceipt returned nil receipt")
	}

	if receipt.TransactionHash().String() != testTxHash.String() {
		t.Fatalf("TxReceipt returned wrong transaction hash: %v", receipt.TransactionHash().String())
	}

	t.Logf("TxReceipt: %+v", receipt.TransactionHash().String())
}

//func TestLogs(t *testing.T) {
//	ctx := context.Background()
//	topics := [][]common.Hash{
//		{common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")},
//	}
//
//	logs, err := cl.Logs(ctx, testAddress, topics, big.NewInt(1234567))
//	if err != nil {
//		t.Fatalf("Logs failed: %v", err)
//	}
//
//	if logs == nil {
//		t.Fatal("Logs returned nil logs")
//	}
//
//	t.Logf("Logs: %+v", logs)
//}
