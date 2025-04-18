package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/s4bb4t/forefinger/pkg/models"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const (
	testContractHex = "0xe6313d1776E4043D906D5B7221BE70CF470F5e87"
	testAddressHex  = "0xEE2213567A282c1e489Cfa4242B06fEebd087203"
	testBlockHash   = "0xdf29aafca34c510304dffd0182ea648afe84efd463f9752bfb35de71d3423b37"
	testTxHashHex   = "0xc48d5e8a0f2746cd885d9246fbfa083e5bd435c607ddbacd4680863ab66deb84"
	latest          = "latest"
	blockNumber     = 22281939
)

var (
	testContract = common.HexToAddress(testContractHex)
	testAddress  = common.HexToAddress(testAddressHex)
	testHash     = common.HexToHash(testBlockHash)
	testTxHash   = common.HexToHash(testTxHashHex)
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

func TestBlockNumber(t *testing.T) {
	ctx := context.Background()

	bn, err := cl.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("BlockNumber failed: %v", err)
	}

	if bn == nil {
		t.Fatal("BlockNumber returned nil receipt")
	}

	t.Logf("BlockNumber: %+v", bn.String())
}

func TestCallContract(t *testing.T) {
	ctx := context.Background()

	erc721ABI := `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

	abiObj, err := abi.JSON(strings.NewReader(erc721ABI))
	if err != nil {
		t.Fatalf("Ошибка при парсинге ABI: %v", err)
	}

	data, err := abiObj.Pack("balanceOf", testAddress)
	if err != nil {
		t.Fatalf("Ошибка при упаковке данных: %v", err)
	}

	msg := models.NewCallMsg().To(testContract).Data(data)
	result, err := cl.CallContract(ctx, msg, "latest")
	if err != nil {
		t.Logf("CallContract завершился с ошибкой: %v", err)
		return
	}

	if result == nil {
		t.Fatal("CallContract вернул nil результат")
	}

	fmt.Println(result)

	var balance *big.Int
	err = abiObj.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		t.Logf("Ошибка при распаковке результата: %v", err)
	} else {
		t.Logf("Баланс токенов: %v", balance.String())
	}
}
