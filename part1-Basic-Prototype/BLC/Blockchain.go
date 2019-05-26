package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block //存储有序的区块
	Tip []byte //最新的区块Hash
	DB  *bolt.DB
}

//增加区块到区块链里面
//func (blc *Blockchain) AddBlockToBlcokchain(data string, height int64, preHash []byte) {
//	//创建新区快
//	newBlock := NewBlock(data, height, preHash)
//	//往链里边添加区块
//	blc.Blocks = append(blc.Blocks, newBlock)
//
//}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//尝试取表对象
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic("created table failed: ", err)
		}

		if b == nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data......")
			//将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic("存储失败: ", err)
			}

			//存储最新的区块的Hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic("存储Hash失败: ", err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})

	//返回区块链对象
	return &Blockchain{blockHash, db}
}
