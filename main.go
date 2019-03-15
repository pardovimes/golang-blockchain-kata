package main

import (
	"blockgo/pkg/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()
	fmt.Println(bc.GetLastBlock())
	for index := 1; index < 50; index++ {
		b := blockchain.Block{
			Index: index,
			Data:  "Bitcoin ganado :D"}
		bc.Push(b)
	}
}
