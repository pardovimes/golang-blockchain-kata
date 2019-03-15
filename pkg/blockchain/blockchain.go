package blockchain

import (
	"fmt"
	"strings"
)

// Blockchain ...
type Blockchain struct {
	Chain      []Block
	difficulty int
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	b := createGenesisBlock()
	bc.Chain = append(bc.Chain, b)
	bc.difficulty = 4
	return bc
}

func createGenesisBlock() Block {
	b := Block{Index: 0, Data: "Genesis"}
	return b
}

func (bc Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc Blockchain) Push(b Block) {
	b.previousHash = bc.GetLastBlock().hash
	b.hash = b.CalculateHash()
	bc.Mine(b)
	bc.Chain = append(bc.Chain, b)
}

func (bc Blockchain) Mine(b Block) {
	for b.hash[:bc.difficulty] != strings.Repeat("0", bc.difficulty) {
		b.nonce++
		b.hash = b.CalculateHash()
	}
	fmt.Println("Minado!")
}
