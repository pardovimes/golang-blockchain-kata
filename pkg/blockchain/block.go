package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block ...
type Block struct {
	Index        int
	timestamp    time.Time
	Data         string
	hash         string
	previousHash string
	nonce        uint64
}

func (b Block) CalculateHash() string {
	hasher := sha256.New()
	stringToHash := strconv.Itoa(b.Index) +
		b.previousHash +
		b.timestamp.String() +
		strconv.Itoa(int(b.nonce))
	hasher.Write([]byte(stringToHash))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}
