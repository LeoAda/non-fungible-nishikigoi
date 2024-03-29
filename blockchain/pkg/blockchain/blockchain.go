package blockchain

// Blockchain object with list of blocks and index of the last block.
type Blockchain struct {
	blocks []*Block //list of blocks
	index  int64    //index of the last block
}

//Get the list of blocks.
func (bc *Blockchain) Blocks() []*Block {
	return bc.blocks
}

//Get index of the last block.
func (bc *Blockchain) Index() int64 {
	return bc.index
}

// Add a new block.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[bc.index]
	bc.index++
	newBlock := NewBlock(bc.index, prevBlock.blockHash, data)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) AddBlockObject(b *Block) {
	bc.index++
	bc.blocks = append(bc.blocks, b)
}

// Create a new genesis block.
func NewGenesisBlock() *Block {
	return NewBlock(0, "", "Genesis Block")
}

// Create a blockchain.
func NewBlockchain(b *Block) *Blockchain {
	if b == nil {
		return &Blockchain{[]*Block{NewGenesisBlock()}, 0}
	} else {
		return &Blockchain{[]*Block{b}, 0}
	}
}
func NewBlockchainList(b []*Block) *Blockchain {
	return &Blockchain{b, int64(cap(b)) - 1}
}
