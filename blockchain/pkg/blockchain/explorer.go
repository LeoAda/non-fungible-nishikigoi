package blockchain

/*
func (bi *BlockIterator) hasNext() bool {
	return bi.currentBlock.ParentHash() != ""

}

func (bi *BlockIterator) getNext() *Block {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
*/

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
