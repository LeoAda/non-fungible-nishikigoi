package storage

import "leoada/blockchain/pkg/blockchain"

//Interface for database use
type Database interface {
	GetBlock(blockHash string) *blockchain.Block
	GetLastBlock() string
	GetBlockchain() *blockchain.Blockchain
	AddBlock(b *blockchain.Block) error
}
