package blockchain

import (
	"blockchain/websocket"
	"bytes"
	"encoding/json"
	"fmt"
)

var blockchain []Block

//Init initialisation de la block chain
func Init() {
	if len(blockchain) == 0 {
		blockchain = append(blockchain, generateFirstBlock())
	}
}
func replaceChain(newBlockChain []Block) {
	if isValidChain(newBlockChain) && (len(newBlockChain) > len(blockchain)) {
		fmt.Println("Received blockchain is valid. Replacing current blockchain with received blockchain")
		blockchain = newBlockChain
		//		broadcast(responseLatestMsg())
	} else {
		fmt.Println("Received blockchain invalid")
	}
}

func isValidChain(blockchainToValidate []Block) bool {
	tmp1, _ := json.Marshal(blockchainToValidate[0])
	tmp2, _ := json.Marshal(generateFirstBlock())
	if !bytes.Equal(tmp1, tmp2) {
		return false
	}
	tempBlocks := blockchainToValidate[0]
	for i := 1; i < len(blockchainToValidate); i++ {
		if validateNewBlock(blockchainToValidate[i], tempBlocks) {
			tempBlocks = blockchainToValidate[i]
		} else {
			return false
		}
	}
	return true
}

//AddBlock Ajoute un block a la blockChain
func AddBlock(newBlock Block) {
	if validateNewBlock(newBlock, getLatestBlock()) {
		blockchain = append(blockchain, newBlock)
	}
}

//AddDataBlock Ajoute un block a la blockChain
func AddDataBlock(data string) {
	// Creation d'un blockl
	newBlock := GenerateNextBlock(getLatestBlock(), data)
	AddBlock(newBlock)
	b, _ := json.Marshal(blockchain)
	PeerToPeer.SendMessageToHub(string(b))
}

func getLatestBlock() Block {
	return blockchain[len(blockchain)-1]
}
