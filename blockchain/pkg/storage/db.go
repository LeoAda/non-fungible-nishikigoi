package storage

import "leoada/blockchain/pkg/blockchain"

type database interface {
	GetBlock(blockHash string) blockchain.Block
	GetLastBlock() string
	GetBlockchain() blockchain.Blockchain
	AddBlock(b blockchain.Block) bool
}
