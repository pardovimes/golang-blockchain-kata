package main

import (
	"blockgo/pkg/blockchain"
	"fmt"
	"strconv"
	"time"
)

func main() {
	bc := blockchain.NewBlockchain()
	for index := 1; index < 10; index++ {
		b := blockchain.Block{
			Index:     index,
			Data:      "amount: " + strconv.Itoa(50),
			Timestamp: time.Now()}
		bc.Push(b)
	}
	fmt.Println(bc.IsValid())
}
