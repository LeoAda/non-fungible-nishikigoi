package storage

import "leoada/blockchain/pkg/blockchain"

//Interface for database use
type Database interface {
	GetBlock(blockHash string) (*blockchain.Block, error)
	GetLastBlock() string
	GetBlockchain() (*blockchain.Blockchain, error)
	AddBlock(b *blockchain.Block) error
}
