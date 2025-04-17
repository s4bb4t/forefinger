# Forefinger

Forefinger is a high-performance Go library for interacting with Ethereum JSON-RPC API. It simplifies working with
Ethereum nodes by providing an optimized connection pool, efficient memory management, and a streamlined interface. The
library reduces the complexity of handling RPC requests, parsing responses, and managing concurrent connections,
allowing developers to focus on their application logic rather than infrastructure details.

## Features

- Efficient connection pool management
- Batch calls for performance optimization
- Support for all standard Ethereum JSON-RPC methods
- Memory-efficient data models with compressed storage for rarely used fields
- Log and event filtering capabilities
- Developer-friendly API for Go applications

## Installation

```bash
go get -u github.com/s4bb4t/forefinger
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/s4bb4t/forefinger/pkg/client"
	"math/big"
)

func main() {
	// Create a client with Ethereum node URL and connection pool size
	c, err := client.NewClient("https://mainnet.infura.io/v3/YOUR-API-KEY", 5)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Get the latest block
	block, err := c.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// Print the block number
	fmt.Printf("Latest block: %s\n", block.Number())
}
```

## Core Capabilities

### Client Initialization

```go
// Create a client with a pool of 5 connections
client, err := client.NewClient("https://mainnet.infura.io/v3/YOUR-API-KEY", 5)
if err != nil {
// handle error
}
defer client.Close()
```

### Block and Transaction Queries

```go
// Get block by number
block, err := client.BlockByNumber(ctx, big.NewInt(14000000))

// Get block by hash
blockHash := common.HexToHash("0x...")
block, err := client.BlockByHash(ctx, blockHash)

// Get transaction by hash
txHash := common.HexToHash("0x...")
tx, err := client.TxByHash(ctx, txHash)

// Get transaction receipt
receipt, err := client.TxReceipt(ctx, txHash)
```

### Account and Smart Contract Operations

```go
// Get account balance
address := common.HexToAddress("0x...")
balance, err := client.Balance(ctx, address, nil)

// Read smart contract code
code, err := client.Code(ctx, address, nil)

// Call smart contract method
data := []byte{...} // ABI-encoded call data
result, err := client.CallContract(ctx, address, data, nil)

// Estimate gas for a transaction
gas, err := client.EstimateGas(ctx, data, nil)
```

### Filters and Logs

```go
// Create a new filter
filter := models.NewFilter().
FromBlock("latest"). // Starting block
ToBlock(big.NewInt(14000000)). // End block
Address("0x123...").           // Contract address filter
Topic("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // Event signature for Transfer

// Apply filter to get logs
logs, err := client.Logs(ctx, filter)

// Create and use node-side filter
filterId, err := client.NewFilter(ctx, filter)
if err != nil {
// handle error
}

// Get filter changes
changes, err := client.FilterChanges(ctx, filterId)

// Remove filter after use
client.UninstallFilter(ctx, filterId)
```

## Batch Requests

Forefinger supports two types of batch requests for performance optimization:

### BatchCall

Executes multiple requests of the same type in a single HTTP call:

```go
// Create results slice
results := make([]*big.Int, 10)

// Create arguments slice
args := make([][]any, 10)
for i := 0; i < 10; i++ {
args[i] = []any{common.HexToAddress(fmt.Sprintf("0x%x", i)), "latest"}
}

// Execute batch balance request
err, errs := client.BatchCall(
ctx,
5, // batch size
methods.Balance,
args,
&results,
)

// Check errors and process results
```

### SequenceBatchCall

Executes a sequence of different request types:

```go
// Create sequence of requests
sequence := methods.Sequence{
{
Method: methods.BlockByNumber,
Args:   []any{nil, true},
Result: &block,
},
{
Method: methods.Balance,
Args:   []any{common.HexToAddress("0x..."), "latest"},
Result: &balance,
},
}

// Execute batch request
err, errs := client.SequenceBatchCall(ctx, 2, &sequence)
```

## Data Model Structure

Forefinger uses an efficient internal structure for data models, separating commonly and rarely used fields:

```go
// Get block information
block, _ := client.BlockByNumber(ctx, nil)

// Access frequently used fields (stored directly in the structure)
blockNumber := block.Number() // *big.Int
timestamp := block.Timestamp() // *big.Int
txs := block.Transactions() // Transactions

// Access additional fields (stored in compressed format)
hash, _ := block.Hash() // common.Hash
miner, _ := block.Miner() // common.Address
difficulty, _ := block.Difficulty() // *big.Int
```

## Concurrent Processing

The library ensures safe concurrent operation using a connection pool:

```go
// Create client with pool of 10 connections
client, _ := client.NewClient("https://ethereum.node.url", 10)

// Execute parallel requests
var wg sync.WaitGroup
for i := 0; i < 20; i++ {
wg.Add(1)
go func (blockNum int64) {
defer wg.Done()
block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNum))
// process block
}(int64(15000000 + i))
}
wg.Wait()
```

## License

[MIT License](LICENSE)