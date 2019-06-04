package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
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
		fmt.Println("Txs:")
		for _, tx := range block.Txs {
			fmt.Printf("%x\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("%x\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
				fmt.Printf("%s\n", in.ScriptSig)
			}
			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Println(out.Value)
				fmt.Println(out.ScriptPubKey)
			}
		}
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println("--------------------")
		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

	}

}

func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlcokchain(txs []*Transaction) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//2.获取最新区块,并反序列化
			block := DeserializeBlock(b.Get(blc.Tip))

			//3.将区块序列化并且存储到数据库中
			newBlock := NewBlock(txs, block.Height+1, block.Hash)
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
func CreateBlockchainWithGenesisBlock(address string) *Blockchain {

	if DBExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块...")

	//创世区块不存在
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var genesisHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//创建表对象
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic("created table failed: ", err)
		}

		if b != nil {

			//创建了一个coinbase transaction
			txCoinbase := NewCoinbaseTransaction(address)
			//创建创世区块
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
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
			genesisHash = genesisBlock.Hash
		}

		return nil
	})
	return &Blockchain{genesisHash, db}
}

//返回Blockchain对象
func BlockchainObject() *Blockchain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//读取最新区块的哈希
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

//挖掘新的区块
func (blockchain *Blockchain) MineNewBlock(from []string, to []string, amount []string) {

	//fmt.Println(from)
	//fmt.Println(to)
	//fmt.Println(amount)
	//1.建立一笔交易
	value, _ := strconv.Atoi(amount[0])

	tx := NewSimpleTransaciton(from[0], to[0], value)

	//1.通过相关算法建立交易transaction数组
	var txs []*Transaction
	txs = append(txs, tx)

	var block *Block

	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte("l"))

			block = DeserializeBlock(b.Get(hash))
		}
		return nil
	})

	//2.挖新的区块
	block = NewBlock(txs, block.Height+1, block.Hash)

	//3.存进数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			b.Put(block.Hash, block.Serialize())
			b.Put([]byte("l"), block.Hash)
			blockchain.Tip = block.Hash

		}
		return nil
	})

}
