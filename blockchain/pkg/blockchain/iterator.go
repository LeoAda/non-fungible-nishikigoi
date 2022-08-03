package blockchain

type Iterator struct {
	currentBlock *Block
	bc           *Blockchain
}

func (bc *Blockchain) createIterator() *Iterator {
	return &Iterator{
		currentBlock: bc.LastBlock(),
		bc:           bc,
	}
}

func (i *Iterator) resetIterator() {
	i.currentBlock = i.bc.LastBlock()
}

func (i *Iterator) Next() *Block {
	if i.hasNext() {
		return nil
	}
	i.currentBlock = i.bc.FindBlock(i.currentBlock.ParentHash())
	return i.currentBlock
}

func (i *Iterator) hasNext() bool {
	return i.currentBlock.ParentHash() != ""
}
