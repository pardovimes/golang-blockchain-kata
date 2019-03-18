package main

import (
	"blockgo/pkg/blockchain"
	"fmt"
	"strconv"
	"time"
)

func main() {
	bc := blockchain.NewBlockchain()
	fmt.Println(bc.GetLastBlock())
	for index := 1; index < 50; index++ {
		b := blockchain.Block{
			Index:     index,
			Data:      "amount: " + strconv.Itoa(50),
			Timestamp: time.Now()}
		bc.Push(b)
	}
}
