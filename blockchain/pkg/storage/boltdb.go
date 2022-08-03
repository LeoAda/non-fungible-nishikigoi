package storage

import (
	"errors"
	"fmt"
	"leoada/blockchain/pkg/blockchain"

	"github.com/boltdb/bolt"
)

const dbName string = "blockchain.db"

//Implementation of the database interface for BoltDb
type Boltdb struct {
	boltDbInstance *bolt.DB
}

//Getter for the boltdb instance, init it if it's nil
func (b *Boltdb) getDb() *bolt.DB {
	if b.boltDbInstance == nil {
		var err error
		b.boltDbInstance, err = bolt.Open(dbName, 0600, nil)
		if err != nil {
			panic("Pb db")
		}
	}
	return b.boltDbInstance
}

//Add a block to the database, as key : hashBlock and value : serialize value of the block
func (b *Boltdb) AddBlock(block *blockchain.Block) bool {
	db := b.getDb()
	db.Update(func(tx *bolt.Tx) error {
		var err error
		bkt := tx.Bucket([]byte("blocksBucket"))
		if bkt == nil {
			bkt, err = tx.CreateBucket([]byte("blocksBucket"))
			if err != nil {
				panic("Pb db")
			}
			fmt.Println("create")
		}
		serialization, err := SerializeBlock(block)
		blockHash := block.BlockHash()
		err = bkt.Put([]byte(blockHash), []byte(serialization))
		err = bkt.Put([]byte("l"), []byte(blockHash))
		return err
	})

	return true
}

//Get the hash of the last block added to the db, it's at the key "l"
func (b *Boltdb) GetLastBlock() string {
	var blockHash string
	db := b.getDb()
	db.View(func(tx *bolt.Tx) error {
		var err error
		bkt := tx.Bucket([]byte("blocksBucket"))
		if bkt == nil {
			return errors.New("no value in db")
		}

		blockHash = string(bkt.Get([]byte("l")))
		return err
	})
	return blockHash
}

//Get block object from a hash string
func (b *Boltdb) GetBlock(s string) *blockchain.Block {
	var blockHash string
	db := b.getDb()
	db.View(func(tx *bolt.Tx) error {
		var err error
		bkt := tx.Bucket([]byte("blocksBucket"))
		if bkt == nil {
			return errors.New("no value in db")
		}
		blockHash = string(bkt.Get([]byte(s)))
		return err
	})
	block, _ := DeserializeBlock(blockHash)
	return block
}

func (b *Boltdb) GetBlockchain() *blockchain.Blockchain {
	blockHash := b.GetLastBlock()
	block := b.GetBlock(blockHash)
	size := block.BlockNumber()
	blockList := make([]*blockchain.Block, size+1)
	blockList[block.BlockNumber()] = block
	for block.ParentHash() != "" {
		block = b.GetBlock(block.ParentHash())
		blockList[block.BlockNumber()] = block
	}
	return blockchain.NewBlockchainList(blockList)
}
