package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block //存储有序的区块
	Tip []byte //最新的区块Hash
	DB  *bolt.DB
}

//遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	blockchainIterator := blc.Iterator()

	for {
		block := blockchainIterator.Next()

		fmt.Println("新区快：")
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PreBlockHash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()
		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

	}

}

func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlcokchain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//2.获取最新区块,并反序列化
			block := DeserializeBlock(b.Get(blc.Tip))

			//3.将区块序列化并且存储到数据库中
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4.更新数据库中l对应的hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5.更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	//判断创世区块是否存在
	if dbExists() {
		fmt.Println("创世区块已经存在")
		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}

		var blockChain *Blockchain
		//如果创世区块存在，数据库必然也存在，取出数据库最新区块的hash和db返回
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			hash := b.Get([]byte("l"))
			blockChain = &Blockchain{hash, db}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockChain
	}

	//创世区块不存在
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//尝试获取表对象，不存在就创建
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//创建表对象
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic("created table failed: ", err)
			}
		} else {
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
