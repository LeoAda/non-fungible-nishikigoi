package main

import (
	"fmt"
	"leoada/blockchain/pkg/blockchain"
	"leoada/blockchain/pkg/storage"
)

//TODO : add bpm to blockchain
//TODO : Add pow
//TODO : Handle the data separator with serilizer and deserializer
//TODO : Handle error globally
func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	var db storage.Database = &storage.Boltdb{}
	for _, block := range bc.Blocks() {
		fmt.Printf("index: %x\n", block.BlockNumber())
		fmt.Printf("Prev. hash: %s\n", block.ParentHash())
		fmt.Printf("Data: %s\n", block.Data())
		fmt.Printf("Hash: %s\n", block.BlockHash())
		fmt.Println()
		db.AddBlock(block)
	}
	lastblockhash := db.GetLastBlock()
	fmt.Printf("Last block Hash: %s\n", lastblockhash)
	lastblock := db.GetBlock(lastblockhash)
	fmt.Printf("Last index: %x\n", lastblock.BlockNumber())
	fmt.Printf("last Prev. hash: %s\n", lastblock.ParentHash())
	fmt.Printf("last Data: %s\n", lastblock.Data())
	fmt.Printf("Last Hash: %s\n", lastblock.BlockHash())
	fmt.Println()
	bcc := db.GetBlockchain()
	for _, block := range bcc.Blocks() {
		fmt.Printf("index: %x\n", block.BlockNumber())
		fmt.Printf("Prev. hash: %s\n", block.ParentHash())
		fmt.Printf("Data: %s\n", block.Data())
		fmt.Printf("Hash: %s\n", block.BlockHash())
		fmt.Println()
		db.AddBlock(block)
	}
}
