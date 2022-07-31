package storage

import (
	"errors"
	"fmt"
	"leoada/blockchain/pkg/blockchain"

	"github.com/boltdb/bolt"
)

var dbInstance *bolt.DB

func getDb() *bolt.DB {
	if dbInstance == nil {
		var err error
		dbInstance, err = bolt.Open("blockchain.db", 0600, nil)
		if err != nil {
			panic("Pb db")
		}
		return dbInstance
	} else {
		return dbInstance
	}

}

func AddBlock(b *blockchain.Block) bool {
	db := getDb()
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
		serialization, err := SerializeBlock(b)
		blockHash := b.BlockHash()
		err = bkt.Put([]byte(blockHash), []byte(serialization))
		err = bkt.Put([]byte("l"), []byte(blockHash))
		return err
	})

	return true
}

func GetLastBlock() string {
	var blockHash string
	db := getDb()
	db.View(func(tx *bolt.Tx) error {
		var err error
		bkt := tx.Bucket([]byte("blocksBucket"))
		if bkt == nil {
			return errors.New("No value in db")
		}

		blockHash = string(bkt.Get([]byte("l")))
		return err
	})
	return blockHash
}

func GetBlock(s string) *blockchain.Block {
	var blockHash string
	db := getDb()
	db.View(func(tx *bolt.Tx) error {
		var err error
		bkt := tx.Bucket([]byte("blocksBucket"))
		if bkt == nil {
			return errors.New("No value in db")
		}
		blockHash = string(bkt.Get([]byte(s)))
		return err
	})
	block, _ := DeserializeBlock(blockHash)
	return block
}

// Key => value
// hash => serialize
// l => last hash
