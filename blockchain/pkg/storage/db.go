package storage

import (
	"errors"
	"fmt"
	"leoada/blockchain/pkg/blockchain"
	"strconv"
	"strings"
)

//SerializeBlock serialize a block to a string for the database. It use ";" as a separator.
//The order is : blockNumber;timestamp;parentHash;data;blockHash.
//If getdata contains an ";", it return an empty string and an error.
func SerializeBlock(b *blockchain.Block) (string, error) {
	//check if getData contain ;, else return error
	if strings.Contains(b.Data(), ";") {
		return "", errors.New("data contain ;")
	}
	return fmt.Sprintf("%d;%d;%s;%s;%s", b.BlockNumber(), b.Timestamp(), b.ParentHash(), b.Data(), b.BlockHash()), nil
}

//DeserializeBlock deserialize a block from a string. It use ";" as a separator.
//The order is : blockNumber;timestamp;parentHash;data;blockHash.
// if blocknumber or timestamp is not a string, it return a nil block and an error
// if the list have not a len of 5, it return a nil block and an erro
func DeserializeBlock(s string) (*blockchain.Block, error) {
	//split the string with ;
	list := strings.Split(s, ";")
	if len(list) != 5 {
		return nil, errors.New("list have not a len of 5")
	}
	//convert the string to int64
	blockNumber, err := strconv.ParseInt(list[0], 10, 64)
	if err != nil {
		return nil, err
	}
	timestamp, err := strconv.ParseInt(list[1], 10, 64)
	if err != nil {
		return nil, err
	}
	//create a new block with the list
	parentHash := list[2]
	data := list[3]
	blockHash := list[4]
	return blockchain.NewBlockAll(blockNumber, timestamp, parentHash, data, blockHash), nil
}
