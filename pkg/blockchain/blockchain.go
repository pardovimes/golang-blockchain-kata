package blockchain

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
	b := Block{Index: 0, data: "Genesis"}
	return b
}

func (bc Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}
