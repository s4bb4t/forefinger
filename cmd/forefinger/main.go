package main

import (
	"context"
	"fmt"
	"github.com/s4bb4t/forefinger/pkg/client"
	"math/big"
)

func main() {
	cl, err := client.NewClient("http://10.255.13.100:8545", 10)
	if err != nil {
		fmt.Println(err)
	}
	defer cl.Close()

	block, err := cl.BlockByNumber(context.Background(), big.NewInt(22000000))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(block.Number(), len(block.Transactions()))
}
