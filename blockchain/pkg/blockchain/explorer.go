package blockchain

func (bc *Blockchain) FindBlock(blockHash string) *Block {
	for _, n := range bc.blocks {
		if n.BlockHash() == blockHash {
			return n
		}
	}
	return nil
}

func (bc *Blockchain) LastBlock() *Block {
	for _, n := range bc.blocks {
		if n.BlockNumber() == bc.Index() {
			return n
		}
	}
	return nil
}

func (bc *Blockchain) BlockExist(b *Block) bool {
	for _, n := range bc.blocks {
		if n.BlockHash() == b.BlockHash() {
			return true
		}
	}
	return false
}
