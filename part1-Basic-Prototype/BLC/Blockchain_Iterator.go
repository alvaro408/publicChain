package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchainIterator *BlockchainIterator) Next() *Block {

	var block *Block

	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			block = DeserializeBlock(b.Get([]byte(blockchainIterator.CurrentHash)))
			//更新迭代器的hash值
			blockchainIterator.CurrentHash = block.PreBlockHash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}
