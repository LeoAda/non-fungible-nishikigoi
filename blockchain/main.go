package main

import (
	"fmt"
	"leoada/blockchain/pkg/blockchain"
)

//TODO : add bpm to blockchain
//TODO : Add pow
//TODO : change variable of block to private (private = all package file)
//TODO : Handle the data separator with serilizer and deserializer

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks() {
		fmt.Printf("index: %x\n", block.BlockNumber())
		fmt.Printf("Prev. hash: %s\n", block.ParentHash())
		fmt.Printf("Data: %s\n", block.Data())
		fmt.Printf("Hash: %s\n", block.BlockHash())
		fmt.Println()
	}

}
