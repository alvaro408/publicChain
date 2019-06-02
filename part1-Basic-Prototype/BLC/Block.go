package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块的hash
	PreBlockHash []byte
	//3. 交易数据
	Txs []*Transaction
	//4. 时间戳
	Timestamp int64
	//5. hash
	Hash []byte
	//6. Nonce
	Nonce int64
}

//需要将Txs转换成字节数组
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

//将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	//创建缓冲区
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

//反序列化
func DeserializeBlock(blockBytes []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//1. 创建新的区块
func NewBlock(txs []*Transaction, height int64, preBlockHash []byte) *Block {
	//1. 生成区块
	block := &Block{height, preBlockHash, txs, time.Now().Unix(), nil, 0}
	//2. 调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//2. 生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {

	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

//func (block *Block) SetHash() {
//	//1. Height转换为字节数组
//	heightBytes := IntToHex(block.Height)
//	//2. 时间戳转换为字节数组
//	timeString := strconv.FormatInt(block.Timestamp, 2)
//	timeByte := []byte(timeString)
//	//3. 拼接所有属性
//	blockBytes := bytes.Join([][]byte{heightBytes, block.PreBlockHash, block.Data, timeByte, block.Hash}, []byte{})
//	//4. 生成Hash
//	hash := sha256.Sum256(blockBytes)
//	block.Hash = hash[:]
//}
