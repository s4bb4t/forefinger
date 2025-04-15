package main

import (
	"context"
	"fmt"
	"github.com/s4bb4t/forefinger/internal/models"
	"github.com/s4bb4t/forefinger/pkg/client"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"math/big"
)

func main() {
	cl, err := client.NewClient("http://10.255.13.100:8545", 10)
	if err != nil {
		fmt.Println(err)
	}
	defer cl.Close()

	var block models.Block
	if err := cl.Call(context.Background(), &block, methods.BlockByNumber, big.NewInt(22000000), true); err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(block.Transactions()))

	for i := 0; i < len(block.Transactions()); i++ {
		fmt.Println()
		fmt.Println(block.Transactions()[i].BlockNumber())
		fmt.Println(block.Transactions()[i].Hash())
		fmt.Println(block.Transactions()[i].To())
		fmt.Println(block.Transactions()[i].From())
		fmt.Println(block.Transactions()[i].Input())
		fmt.Println(block.Transactions()[i].Value())
		fmt.Println(block.Transactions()[i].V())
		fmt.Println(block.Transactions()[i].S())
		fmt.Println(block.Transactions()[i].R())
	}
}
