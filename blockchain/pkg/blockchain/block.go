package blockchain

import (
	"errors"
	"leoada/blockchain/pkg/security"
	"strconv"
	"strings"
	"time"
)

//Block object from a blockchain.
// In order : blockNumber, timestamp, parentHash, data, blockHash
type Block struct {
	blockNumber int64  //the number of the block in the blockchain
	timestamp   int64  //the time the block was created
	parentHash  string //the hash of the previous block
	data        string //the data of the block
	blockHash   string //the hash of the block
}

// Create a new block with the given data, parent hash and block number. it use NewBlockAll and fix the timestamp.
func NewBlock(currentBlockchainindex int64, parentHash string, data string) (*Block, error) {
	return NewBlockAll(currentBlockchainindex, time.Now().Unix(), parentHash, data, "")
}

// Create a new block with the given data, parent hash, block number and timestamp.
func NewBlockAll(currentBlockchainindex int64, timestamp int64, parentHash string, data string, blockHash string) (*Block, error) {
	var block *Block
	if blockHash == "" {
		block = &Block{currentBlockchainindex, timestamp, parentHash, data, ""}
		block.setHash()
	} else {
		block = &Block{currentBlockchainindex, timestamp, parentHash, data, blockHash}
		if blockHash != block.calculateBlockHash() {
			return nil, errors.New("NewBlockAll : Wrong hash")
		}
	}
	return block, nil
}

// Get the number of the block.
func (b *Block) BlockNumber() int64 {
	return b.blockNumber
}

// Get the time of creation of the block.
func (b *Block) Timestamp() int64 {
	return b.timestamp
}

// Get the hash of the previous block.
func (b *Block) ParentHash() string {
	return b.parentHash
}

// Get the data of the block.
func (b *Block) Data() string {
	return b.data
}

// Get the hash of the block.
func (b *Block) BlockHash() string {
	return b.blockHash
}

//Set the hash of the block with his parameters. The hash is calculated with sha256.
func (b *Block) setHash() {
	b.blockHash = b.calculateBlockHash()
}

//Calculate hash of a block and return it
func (b *Block) calculateBlockHash() string {
	timestamp := strconv.FormatInt(b.timestamp, 10)                      //convert timestamp to string
	blockNumber := strconv.FormatInt(b.blockNumber, 10)                  //convert blockNumber to string
	headerList := []string{blockNumber, timestamp, b.parentHash, b.data} //create header list of strings
	header := strings.Join(headerList, "")                               //join header list to strings
	hash := security.HashString(header)                                  //calculate hash
	return hash
}
