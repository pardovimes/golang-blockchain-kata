package main

import (
	"blockgo/pkg/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()
	fmt.Println(bc.GetLastBlock())
}
