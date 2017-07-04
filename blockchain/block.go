package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Index           int64     `json:"id" bson:"id"`
	Timestamp       time.Time `json:"timestamp" bson:"timestamp"`
	Data            string    `json:"data" bson:"data"`
	HashKey         []byte    `json:"key" bson:"key"`
	PreviousHAshKey []byte    `json:"pKey" bson:"pKey"`
}

func createBlock(index int64, timestamp time.Time, data string, hashKey []byte, previousHAshKey []byte) (createdBlock Block) {
	return Block{index, timestamp, data, hashKey, previousHAshKey}
}

func calculateHash(index int64, previousHash []byte, timestamp time.Time, data string) (computedHashKey []byte) {
	t := timestamp.Format(time.UnixDate)
	key := string(index) + string(previousHash) + string(t) + data
	hashKey := sha256.Sum256([]byte(key))
	return hashKey[:]
}

//GenerateNextBlock Generation du prochain block
func GenerateNextBlock(previousBlock Block, data string) (generatedBlock Block) {
	index := previousBlock.Index + 1
	currentTimestamp := time.Now()
	newHash := calculateHash(index, previousBlock.HashKey, currentTimestamp, data)
	return Block{index, currentTimestamp, data, newHash, previousBlock.HashKey}
}

func generateFirstBlock() (firstBlock Block) {
	index := int64(0)
	currentTimestamp := time.Now()
	firstHash := make([]byte, 1, 1)
	newHash := calculateHash(index, firstHash, currentTimestamp, "")
	return Block{0, currentTimestamp, "", newHash, firstHash}
}
func validateNewBlock(newBlock Block, previousBlock Block) bool {
	if previousBlock.Index+1 != newBlock.Index {

		return false
	}

	if !bytes.Equal(previousBlock.HashKey, newBlock.PreviousHAshKey) {
		return false
	}

	if !bytes.Equal(calculateHash(newBlock.Index, previousBlock.HashKey, newBlock.Timestamp, newBlock.Data), newBlock.HashKey) {

		return false
	}

	return true
}
