package blockchain

// Blockchain object with list of blocks and index of the last block.
type blockchain struct {
	blocks []*Block //list of blocks
	index  int64    //index of the last block
}

//Get the list of blocks.
func (bc *blockchain) Blocks() []*Block {
	return bc.blocks
}

//Get index of the last block.
func (bc *blockchain) Index() int64 {
	return bc.index
}

// Add a new block.
func (bc *blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[bc.index]
	bc.index++
	newBlock := NewBlock(bc.index, prevBlock.blockHash, data)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create a new genesis block.
func NewGenesisBlock() *Block {
	return NewBlock(0, "", "Genesis Block")
}

// Create a blockchain.
func NewBlockchain() *blockchain {
	return &blockchain{[]*Block{NewGenesisBlock()}, 0}
}

//  the last block.
func (bc *blockchain) LastBlock() *Block {
	return bc.blocks[bc.index]
}
